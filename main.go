package main

import (
	"fmt"

	"github.com/minhduccm/go-cleanarch-boilerplate/databases/mongodb"
	"github.com/minhduccm/go-cleanarch-boilerplate/repositories"
	"github.com/minhduccm/go-cleanarch-boilerplate/usecases"
	webservices "github.com/minhduccm/go-cleanarch-boilerplate/webservices"
)

func main() {
	fmt.Println("Starting...")
	mongoDB := mongodb.NewMongoDB()
	bookRepo := repositories.NewBookRepo(mongoDB)
	cateRepo := repositories.NewCateRepo(mongoDB)
	bookUsecase := usecases.NewBookUsecase(bookRepo, cateRepo)
	webService := webservices.NewWebService(bookUsecase)
	webService.Run()
}

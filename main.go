package main

import (
	"fmt"

	"github.com/minhduccm/go-cleanarch-boilerplate/databases/mongodb"
	"github.com/minhduccm/go-cleanarch-boilerplate/repositories"
	"github.com/minhduccm/go-cleanarch-boilerplate/usecases"
	"github.com/minhduccm/go-cleanarch-boilerplate/webframeworks/martini"
	"github.com/minhduccm/go-cleanarch-boilerplate/wshandlers"
)

func main() {
	fmt.Println("Starting...")
	mongoDB := mongodb.NewMongoDB()
	bookRepo := repositories.NewBookRepo(mongoDB)
	cateRepo := repositories.NewCateRepo(mongoDB)
	bookUsecase := usecases.NewBookUsecase(bookRepo, cateRepo)
	webService := wshandlers.NewWebService(bookUsecase)
	webFW := martini.NewMartiniWF(webService)
	webFW.Run()
}

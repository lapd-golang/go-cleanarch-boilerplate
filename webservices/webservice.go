package wshandlers

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/minhduccm/go-cleanarch-boilerplate/entities"
)

type BookUsecase interface {
	GetBooksByCate(int) ([]*entities.Book, error)
	CreateBook(int, string, float64, int) (*entities.Book, error)
	CreateCate(int, string) (*entities.Category, error)
}

type WebService struct {
	BookUsecase BookUsecase
}

func NewWebService(bookUsecase BookUsecase) *WebService {
	return &WebService{
		BookUsecase: bookUsecase,
	}
}

func (ws *WebService) Run() {
	m := martini.Classic()

	m.Post("/categories", ws.CreateCate)
	m.Post("/categories/:cateId/books", ws.CreateBook)
	m.Get("/categories/:cateId/books", ws.ShowBooksByCate)

	m.Run()
	fmt.Println("Listening on port 3000")
}

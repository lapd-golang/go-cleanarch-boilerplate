package wshandlers

import "github.com/minhduccm/go-cleanarch-boilerplate/entities"

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

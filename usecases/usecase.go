package usecases

import "github.com/minhduccm/go-cleanarch-boilerplate/entities"

type BookRepo interface {
	FindBooksByCateId(int) ([]*entities.Book, error)
	Store(*entities.Book) (*entities.Book, error)
}

type CateRepo interface {
	FindCateById(int) (*entities.Category, error)
	Store(*entities.Category) (*entities.Category, error)
}

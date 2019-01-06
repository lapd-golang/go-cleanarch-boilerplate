package usecases

import (
	"github.com/minhduccm/go-cleanarch-boilerplate/entities"
)

type BookUsecase struct {
	BookRepo BookRepo
	CateRepo CateRepo
}

func NewBookUsecase(
	bookRepo BookRepo,
	cateRepo CateRepo,
) *BookUsecase {
	return &BookUsecase{
		BookRepo: bookRepo,
		CateRepo: cateRepo,
	}
}

func (bookUsecase *BookUsecase) GetBooksByCate(cateId int) ([]*entities.Book, error) {
	cate, err := bookUsecase.CateRepo.FindCateById(cateId)
	if err != nil {
		return nil, err
	}
	if cate == nil {
		return []*entities.Book{}, nil
	}
	books, err := bookUsecase.BookRepo.FindBooksByCateId(cateId)
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		book.CateName = cate.Name
	}
	return books, nil
}

func (bookUsecase *BookUsecase) CreateBook(
	bookId int,
	name string,
	price float64,
	cateId int,
) (*entities.Book, error) {
	cate, err := bookUsecase.CateRepo.FindCateById(cateId)
	if err != nil {
		return nil, err
	}
	if cate == nil {
		return nil, nil
	}

	book := &entities.Book{
		Id:        bookId,
		Name:      name,
		Price:     price,
		Available: true,
		CateId:    cate.Id,
		CateName:  cate.Name,
	}
	createdBook, err := bookUsecase.BookRepo.Store(book)
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}

func (bookUsecase *BookUsecase) CreateCate(
	cateId int,
	cateName string,
) (*entities.Category, error) {
	cate := &entities.Category{
		Id:   cateId,
		Name: cateName,
	}
	createdCate, err := bookUsecase.CateRepo.Store(cate)
	if err != nil {
		return nil, err
	}
	return createdCate, nil
}

package usecases

import (
	"testing"

	"github.com/minhduccm/go-cleanarch-boilerplate/entities"
)

type BookRepoFaker struct{}
type CategoryRepoFaker struct{}

func (brf *BookRepoFaker) FindBooksByCateId(
	cateId int,
) ([]*entities.Book, error) {
	book1 := &entities.Book{
		Id:        1,
		Name:      "Book 1",
		Price:     12.5,
		Available: true,
		CateId:    1,
	}
	book2 := &entities.Book{
		Id:        2,
		Name:      "Book 2",
		Price:     30.5,
		Available: true,
		CateId:    1,
	}
	return []*entities.Book{book1, book2}, nil
}

func (brf *BookRepoFaker) Store(
	book *entities.Book,
) (*entities.Book, error) {
	return book, nil
}

func (crf *CategoryRepoFaker) FindCateById(
	cateId int,
) (*entities.Category, error) {
	return &entities.Category{
		Id:   cateId,
		Name: "Category 1",
	}, nil
}

func (crf *CategoryRepoFaker) Store(
	cate *entities.Category,
) (*entities.Category, error) {
	return cate, nil
}

func TestGetBooksByCate(t *testing.T) {
	bookRepoFaker := &BookRepoFaker{}
	categoryRepoFaker := &CategoryRepoFaker{}
	bookUsecase := NewBookUsecase(
		bookRepoFaker,
		categoryRepoFaker,
	)
	categoryId := 1
	books, err := bookUsecase.GetBooksByCate(categoryId)
	if err != nil {
		t.Errorf("GetBooksByCate was incorrect, got error: %v, want list of 2 books to be returned.", err)
		return
	}
	if len(books) != 2 {
		t.Errorf("GetBooksByCate was incorrect, got %d books, want 2 books", len(books))
		return
	}
	cate, _ := categoryRepoFaker.FindCateById(categoryId)
	for _, book := range books {
		if book.CateName != cate.Name {
			t.Errorf("GetBooksByCate was incorrect, got %s, want %s", book.CateName, cate.Name)
			return
		}
		if book.CateId != cate.Id {
			t.Errorf("GetBooksByCate was incorrect, got categoryId = %d, want categoryId = %d", book.CateId, cate.Id)
			return
		}
	}
	t.Log("GetBooksByCate is passed.")
}

func TestCreateBook(t *testing.T) {
	bookRepoFaker := &BookRepoFaker{}
	categoryRepoFaker := &CategoryRepoFaker{}
	bookUsecase := NewBookUsecase(
		bookRepoFaker,
		categoryRepoFaker,
	)
	categoryId := 1
	bookId := 1
	bookName := "Book 1"
	price := 25.5
	createdBook, err := bookUsecase.CreateBook(bookId, bookName, price, categoryId)
	if err != nil {
		t.Errorf("CreateBook was incorrect, got error: %v.", err)
		return
	}
	if createdBook.Id != bookId {
		t.Errorf("CreateBook was incorrect, got book id: %d, want %d", createdBook.Id, bookId)
		return
	}
	if createdBook.Name != bookName {
		t.Errorf("CreateBook was incorrect, got book name: %s, want %s", createdBook.Name, bookName)
		return
	}
	if createdBook.Price != price {
		t.Errorf("CreateBook was incorrect, got book price: %f, want %f", createdBook.Price, price)
		return
	}
	if createdBook.CateId != categoryId {
		t.Errorf("CreateBook was incorrect, got book's categoryId: %d, want %d", createdBook.CateId, categoryId)
		return
	}
	t.Log("CreateBook is passed.")
}

func TestCreateCate(t *testing.T) {
	bookRepoFaker := &BookRepoFaker{}
	categoryRepoFaker := &CategoryRepoFaker{}
	bookUsecase := NewBookUsecase(
		bookRepoFaker,
		categoryRepoFaker,
	)
	categoryId := 1
	categoryName := "Category 1"
	createdCate, err := bookUsecase.CreateCate(categoryId, categoryName)
	if err != nil {
		t.Errorf("CreateBook was incorrect, got error: %v.", err)
		return
	}
	if createdCate.Id != categoryId {
		t.Errorf("CreateBook was incorrect, got category id: %d, want %d", createdCate.Id, categoryId)
		return
	}
	if createdCate.Name != categoryName {
		t.Errorf("CreateBook was incorrect, got category name: %s, want %s", createdCate.Name, categoryName)
		return
	}
	t.Log("CreateCate is passed.")
}

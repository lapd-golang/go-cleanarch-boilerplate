package repositories

import (
	"github.com/minhduccm/go-cleanarch-boilerplate/entities"
)

type BookMapper struct {
	Id        int     `bson:"id"`
	Name      string  `bson:"name"`
	Price     float64 `bson:"price"`
	Available bool    `bson:"available"`
	CateId    int     `bson:"cate_id"`
}

type BookRepo struct {
	Database Database
}

func NewBookRepo(db Database) *BookRepo {
	return &BookRepo{Database: db}
}

func (bookRepo *BookRepo) FindBooksByCateId(
	cateId int,
) ([]*entities.Book, error) {
	rawBooks, err := bookRepo.Database.QueryBooksByCateId(cateId)
	if err != nil {
		return nil, err
	}
	books := make([]*entities.Book, len(rawBooks))
	for i, book := range rawBooks {
		books[i] = &entities.Book{
			Id:        book.Id,
			Name:      book.Name,
			Price:     book.Price,
			Available: book.Available,
			CateId:    book.CateId,
		}
	}
	return books, nil
}

func (bookRepo *BookRepo) Store(
	book *entities.Book,
) (*entities.Book, error) {
	bookMapper := &BookMapper{
		Id:        book.Id,
		Name:      book.Name,
		Price:     book.Price,
		Available: book.Available,
		CateId:    book.CateId,
	}
	err := bookRepo.Database.InsertBook(bookMapper)
	if err != nil {
		return nil, err
	}
	return book, nil
}

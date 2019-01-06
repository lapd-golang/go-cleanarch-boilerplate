package mongodb

import (
	"github.com/minhduccm/go-cleanarch-boilerplate/repositories"
	"gopkg.in/mgo.v2/bson"
)

func (mongoDB *MongoDB) QueryBooksByCateId(cateId int) ([]*repositories.BookMapper, error) {
	sessionCopied := mongoDB.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopied.DB(Database).C(BooksCollection)
	var books []*repositories.BookMapper
	err := collection.Find(bson.M{"cate_id": cateId}).All(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (mongoDB *MongoDB) InsertBook(book *repositories.BookMapper) error {
	sessionCopied := mongoDB.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopied.DB(Database).C(BooksCollection)
	return collection.Insert(book)
}

package mongodb

import (
	"github.com/minhduccm/go-cleanarch-boilerplate/repositories"
	"gopkg.in/mgo.v2/bson"
)

func (mongoDB *MongoDB) QueryCateById(cateId int) (*repositories.CategoryMapper, error) {
	sessionCopy := mongoDB.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(CategoriesCollection)
	var cate repositories.CategoryMapper
	err := collection.Find(bson.M{"id": cateId}).One(&cate)
	if err != nil {
		return nil, err
	}
	return &cate, nil
}

func (mongoDB *MongoDB) InsertCate(cate *repositories.CategoryMapper) error {
	sessionCopy := mongoDB.MongoSession.Copy()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(Database).C(CategoriesCollection)
	return collection.Insert(cate)
}

package mongodb

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const (
	MongoDBHost          = "127.0.0.1:27017"
	Database             = "bookmanagement"
	BooksCollection      = "books"
	CategoriesCollection = "categories"
)

type MongoDB struct {
	MongoSession *mgo.Session
}

func NewMongoDB() *MongoDB {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHost},
		Timeout:  10 * time.Second,
		Database: Database,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, errMongo := mgo.DialWithInfo(mongoDBDialInfo)
	if errMongo != nil {
		log.Fatalf("CreateSession: %s\n", errMongo)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	return &MongoDB{MongoSession: mongoSession}
}

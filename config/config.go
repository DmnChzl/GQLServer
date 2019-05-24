package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// URI : MongoDB URI
const URI = "mongodb://localhost:27017"

// DB : DataBase Name
const DB = "book_store"

// COLL : Collection Name
const COLL = "books"

// Collection : Exported Collection
var Collection *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(DB)
	Collection = db.Collection(COLL)
}

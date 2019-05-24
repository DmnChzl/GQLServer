package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mrdoomy/gqlserver/config"
	"github.com/mrdoomy/gqlserver/models"

	"go.mongodb.org/mongo-driver/bson"
)

func count() (i int) {
	cur, err := config.Collection.Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal("MongoDB Connection Failed...")
	} else {
		fmt.Println("Success To Connect MongoDB !")
	}

	// ForEach
	for cur.Next(context.Background()) {
		i++
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return
}

// LoadFile : Populates DataBase With Some Data (Optional)
func LoadFile(filename string) {
	var books []models.Book

	count := count()

	// Load Values From JSON File To Model
	if count == 0 {
		byteValues, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(byteValues, &books)

		var allBooks []interface{}
		for _, book := range books {
			allBooks = append(allBooks, book)
		}

		_, err = config.Collection.InsertMany(context.Background(), allBooks)
		if err != nil {
			log.Fatal(err)
		}
	}
}

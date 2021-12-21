package queries

import (
	"context"
	"fmt"
	"log"

	"github.com/dmnchzl/gqlserver/config"
	"github.com/dmnchzl/gqlserver/models"
	"github.com/dmnchzl/gqlserver/types"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func forEachBook(cur *mongo.Cursor) (result []models.Book) {
	for cur.Next(context.Background()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	return
}

// GetBooks : Get Many Books Query
var GetBooks = &graphql.Field{
	Type:        graphql.NewList(types.Book),
	Description: "Get Many Books",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: types.IBook,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		filter := params.Args["input"].(map[string]interface{})

		cur, err := config.Collection.Find(context.Background(), filter, nil)
		if err != nil {
			fmt.Println("[GetBooks]: Failed To Find()")
			// log.Fatal(err)
		}

		allBooks := forEachBook(cur)

		return allBooks, nil
	},
}

// GetAllBooks : Get All Books Query
var GetAllBooks = &graphql.Field{
	Type:        graphql.NewList(types.Book),
	Description: "Get All Books",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		cur, err := config.Collection.Find(context.Background(), bson.D{}, nil)
		if err != nil {
			fmt.Println("[GetAllBooks]: Failed To Find()")
			// log.Fatal(err)
		}

		allBooks := forEachBook(cur)

		return allBooks, nil
	},
}

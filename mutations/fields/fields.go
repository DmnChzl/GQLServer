package mutations

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dmnchzl/gqlserver/config"
	"github.com/dmnchzl/gqlserver/models"
	"github.com/dmnchzl/gqlserver/types"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateBook : Create New Book Query
var CreateBook = &graphql.Field{
	Type:        types.Book,
	Description: "Create New Book",
	Args: graphql.FieldConfigArgument{
		"input": &graphql.ArgumentConfig{
			Type: types.IBook,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		input := params.Args["input"].(map[string]interface{})

		// Convert Map To Json
		jsonInput, _ := json.Marshal(input)

		// Convert Json To Struct
		book := models.Book{}
		json.Unmarshal(jsonInput, &book)

		_, err := config.Collection.InsertOne(context.Background(), book)
		if err != nil {
			fmt.Println("[CreateBook]: Failed To InsertOne()")
			// log.Fatal(err)
		}

		return book, nil
	},
}

// UpdateBooks : Update Existing Books Query
var UpdateBooks = &graphql.Field{
	Type:        types.Status,
	Description: "Update Existing Books",
	Args: graphql.FieldConfigArgument{
		"where": &graphql.ArgumentConfig{
			Type: types.IBook,
		},
		"set": &graphql.ArgumentConfig{
			Type: types.IBook,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		filter := params.Args["where"].(map[string]interface{})
		set := params.Args["set"].(map[string]interface{})
		var status models.Status

		res, err := config.Collection.UpdateMany(
			context.Background(),
			filter,
			bson.D{
				{Key: "$set", Value: set},
			})
		if err != nil {
			status.Result = "Failed To UpdateMany..."
			// log.Fatal(err)
		} else {
			status.Result = strconv.FormatInt(res.ModifiedCount, 10) + " Updated."
		}

		return status, nil
	},
}

// DeleteBooks : Delete Existing Books Query
var DeleteBooks = &graphql.Field{
	Type:        types.Status,
	Description: "Delete Existing Books",
	Args: graphql.FieldConfigArgument{
		"where": &graphql.ArgumentConfig{
			Type: types.IBook,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		filter, _ := params.Args["where"].(map[string]interface{})
		var status models.Status

		res, err := config.Collection.DeleteMany(context.Background(), filter)
		if err != nil {
			status.Result = "Failed To DeleteMany..."
			// log.Fatal(err)
		} else {
			status.Result = strconv.FormatInt(res.DeletedCount, 10) + " Deleted."
		}

		return status, nil
	},
}

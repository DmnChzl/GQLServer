package queries

import (
	fields "github.com/mrdoomy/gqlserver/queries/fields"

	"github.com/graphql-go/graphql"
)

// RootQuery : Resolver For Queries
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getAllBooks": fields.GetAllBooks,
		"getBooks":    fields.GetBooks,
	},
})

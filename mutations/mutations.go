package mutations

import (
	fields "github.com/dmnchzl/gqlserver/mutations/fields"

	"github.com/graphql-go/graphql"
)

// RootMutation : Resolver For Mutations
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createBook":  fields.CreateBook,
		"updateBooks": fields.UpdateBooks,
		"deleteBooks": fields.DeleteBooks,
	},
})

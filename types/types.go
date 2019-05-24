package types

import (
	"github.com/graphql-go/graphql"
)

// Book : Book As GraphQL Type
var Book = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"isbn": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"subTitle": &graphql.Field{
			Type: graphql.String,
		},
		"volume": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"date": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"authors": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(Author)),
		},
		"publisher": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"category": &graphql.Field{
			Type: graphql.String,
		},
		"chapters": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"pages": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"price": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Float),
		},
	},
})

// Author : Author As GraphQL Type
var Author = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

// IBook : IBook As GraphQL Type
var IBook = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Book",
	Fields: graphql.InputObjectConfigFieldMap{
		"isbn": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"subTitle": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"volume": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"date": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"authors": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(IAuthor),
		},
		"publisher": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"category": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"chapters": &graphql.InputObjectFieldConfig{
			Type: graphql.NewList(graphql.String),
		},
		"pages": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"price": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
	},
})

// IAuthor : IAuthor As GraphQL Type
var IAuthor = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "Author",
	Fields: graphql.InputObjectConfigFieldMap{
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

// Status : Status As GraphQL Type
var Status = graphql.NewObject(graphql.ObjectConfig{
	Name: "Status",
	Fields: graphql.Fields{
		"result": &graphql.Field{
			Type: graphql.String,
		},
	},
})

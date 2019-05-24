package main

import (
	"log"
	"net/http"

	"github.com/mrdoomy/gqlserver/mutations"
	"github.com/mrdoomy/gqlserver/queries"
	"github.com/mrdoomy/gqlserver/utils"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// PORT : Listening Port
const PORT = "4321"

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.RootQuery,
	Mutation: mutations.RootMutation,
})

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func init() {
	utils.LoadFile("books.json")
}

func main() {
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	http.Handle("/graphql", disableCors(h))
	log.Println("Now Server Is Running On Port " + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	// "github.com/graph-gophers/graphql-go/example/starwars"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema

func init() {
	schema = graphql.MustParseSchema(Schema, &Resolver{})
}

func main() {
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

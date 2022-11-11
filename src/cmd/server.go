package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/LuigiAzevedo/graphql_min_roman_numeral/cmd/graph"
	"github.com/LuigiAzevedo/graphql_min_roman_numeral/cmd/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", server)

	log.Printf("connect to http://localhost:%s/graphql", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

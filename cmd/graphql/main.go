package main

import (
	"log"
	"net/http"
	"os"

	"example.com/pz11-graphql/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("GraphQL server started at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

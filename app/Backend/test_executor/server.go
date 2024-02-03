package main

import (
	"log"
	"net/http"
	"os"
	"test_executor/configdb"
	"test_executor/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8181" 

func main() {

	configdb.ConnectCockroachDB() 

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8180"},
		AllowCredentials: true,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"Settings/configdb"
	"Settings/graph"
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	configdb.ConnectCockroachDB() //db connection + migration

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8080"},
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

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tf63/go-graph-exp/api/graph" // 修正
	"github.com/tf63/go-graph-exp/external"
	"github.com/tf63/go-graph-exp/internal/repository"
	"github.com/tf63/go-graph-exp/internal/resolver" // 修正
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// レイヤードアーキテクチャでAPIを設計
	// DIを入れている
	db, _ := external.ConnectDatabase()
	ntr := repository.NewTodoRepository(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{Tr: ntr}})) // 修正

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

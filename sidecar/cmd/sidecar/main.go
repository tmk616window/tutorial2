package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sidecar/config"
	"sidecar/graph"
	"sidecar/graph/generated"
	"sidecar/infra/db"

	"cloud.google.com/go/storage"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8888"

func main(){
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	if err := db.InitDB(db.URI(cfg.Database), cfg.IsLocal()); err != nil {
		panic(err)
	}

	cfg.GCP.GCSClient, err = storage.NewClient(context.Background())
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}


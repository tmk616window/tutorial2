package main

import (
	"os"
	"sidecar/config"
	"sidecar/infra/db"
	"sidecar/infra/gcs"
	"sidecar/router"
)

const defaultPort = "8888"

func main(){
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	gcsClient, err := gcs.NewClient()
	if err != nil {
		panic(err)
	}

	if err := db.InitDB(db.URI(cfg.Database), cfg.IsLocal()); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := router.ListenAndServe(
		cfg,
		gcsClient,
	); err != nil {
		panic(err)
	}
}


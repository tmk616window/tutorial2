package main

import (
	"fmt"
	"os"
	"sidecar/config"
	"sidecar/infra/db"
	"sidecar/infra/gcs"
	"sidecar/infra/storage"
	"sidecar/router"
)

const defaultPort = "8888"

func main(){
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	var storage storage.StorageCaller
	if cfg.IsLocal() {
		fmt.Println("minio")
	} else {
		storage.StorageInterface, err = gcs.NewClient()
	}
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


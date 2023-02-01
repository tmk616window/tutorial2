package main

import (
	"os"
	"sidecar/config"
	"sidecar/infra/db"
	"sidecar/infra/gcs"
	"sidecar/infra/minio"
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
		storage.StorageInterface, err = minio.NewClient()		
	} else {
		storage.StorageInterface, err = gcs.NewClient()
	}
	if err != nil {
		panic(err)
	}

	db, err := db.InitDB(db.URI(cfg.Database), cfg.IsLocal())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := router.ListenAndServe(
		cfg,
		db,
		storage,
	); err != nil {
		panic(err)
	}
}

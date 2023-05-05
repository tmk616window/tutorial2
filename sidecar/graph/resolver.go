package graph

import (
	"database/sql"
	"sidecar/config"
	"sidecar/infra/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	cfg *config.Cfg
	db *sql.DB
	storage storage.StorageCaller
}

func NewResolver(
	cfg *config.Cfg,
	db *sql.DB,
	storage storage.StorageCaller,
) *Resolver {
	return &Resolver{
		cfg: cfg,
		db: db,
		storage: storage,
	}	
}
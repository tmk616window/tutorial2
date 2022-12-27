package graph

import (
	"sidecar/config"
	"sidecar/infra/gcs"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	cfg *config.Cfg
	gCSService gcs.ClientInterface
}

func NewResolver(
	cfg *config.Cfg,
	gCSService gcs.ClientInterface,
) *Resolver {
	return &Resolver{
		cfg: cfg,
		gCSService: gCSService,
	}	
}
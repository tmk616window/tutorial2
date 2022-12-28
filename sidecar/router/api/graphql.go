package api

import (
	"net/http"

	"sidecar/config"
	"sidecar/graph"
	"sidecar/graph/generated"
	"sidecar/infra/storage"

	"github.com/99designs/gqlgen/graphql/handler"
)

func PostGraphQL(
	cfg *config.Cfg,
	storage storage.StorageCaller,
) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: graph.NewResolver(
			cfg,
			storage,
		)},
	))

	return srv.ServeHTTP
}

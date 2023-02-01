package api

import (
	"database/sql"
	"net/http"

	"sidecar/config"
	"sidecar/graph"
	"sidecar/graph/generated"
	"sidecar/infra/storage"

	"github.com/99designs/gqlgen/graphql/handler"
)

func PostGraphQL(
	cfg *config.Cfg,
	db *sql.DB,
	storage storage.StorageCaller,
) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: graph.NewResolver(
			cfg,
			db,
			storage,
		)},
	))

	return srv.ServeHTTP
}

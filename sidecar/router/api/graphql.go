package api

import (
	"net/http"

	"sidecar/config"
	"sidecar/graph"
	"sidecar/graph/generated"
	"sidecar/infra/gcs"

	"github.com/99designs/gqlgen/graphql/handler"
)

func PostGraphQL(
	cfg *config.Cfg,
	gCSService gcs.ClientInterface,
) http.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: graph.NewResolver(
			cfg,
			gCSService,
		)},
	))

	return srv.ServeHTTP
}

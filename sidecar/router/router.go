package router

import (
	"net/http"
	"time"

	"sidecar/config"
	"sidecar/infra/gcs"
	"sidecar/router/api"

	"github.com/go-chi/chi"
)

func ListenAndServe(
	cfg *config.Cfg,
	gCSService gcs.ClientInterface,
) error {
	router := chi.NewRouter()

	router.Get("/healthcheck", api.GetHealthCheck())
	router.Post("/graphql", api.PostGraphQL(cfg, gCSService))
	// ルーティング
	if cfg.IsLocal() {
		router.Get("/playground", api.GetPlayground("/playground"))
	}

	server := &http.Server{
		Addr:              ":" + cfg.PORT,
		ReadHeaderTimeout: 1 * time.Second,
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       1 * time.Second,
		Handler:           router,
	}

	return server.ListenAndServe()
}

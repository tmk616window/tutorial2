package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
)

func GetPlayground(endpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL", endpoint)
}

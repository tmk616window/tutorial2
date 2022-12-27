package api

import (
	"io"
	"net/http"
)

func GetHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK") //nolint
	}
}

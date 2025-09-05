package httpserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/http/handlers"
)

func NewRouter(h *handlers.HealthHandler) http.Handler {
	r := chi.NewRouter()
	// v0: yalnızca health
	r.Post("/v1/healthz", h.PostHealth)
	return r
}

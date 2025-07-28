package router

import (
	"github.com/go-chi/chi/v5"

	"movilist-api/cmd/api/resource/health"
	"movilist-api/cmd/api/resource/media"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Get)

	r.Route("/v1", func(r chi.Router) {
		mediaAPI := &media.API{}
		r.Get("/media", mediaAPI.List)
		r.Post("/media", mediaAPI.Create)
		r.Get("/media/{id}", mediaAPI.Get)
		r.Put("/media/{id}", mediaAPI.Update)
		r.Delete("/media/{id}", mediaAPI.Delete)
	})

	return r
}

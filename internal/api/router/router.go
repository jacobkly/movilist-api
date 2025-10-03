package router

import (
	"github.com/go-chi/chi/v5"

	"movilist-api/config"
	"movilist-api/internal/api/resource/health"
	"movilist-api/internal/api/resource/movies"
	"movilist-api/pkg/tmdb"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Get)

	c := config.NewTMDB()
	tmdbClient := tmdb.NewClient(c.APIKey)

	r.Route("/v1", func(r chi.Router) {
		movieService := movies.NewService(tmdbClient)
		movieAPI := movies.NewAPI(movieService)

		r.Route("/movies", func(r chi.Router) {
			r.Get("/{id}", movieAPI.GetMovie)
		})
	})

	return r
}

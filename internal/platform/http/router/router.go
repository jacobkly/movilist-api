package router

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"klyvi-api/internal/health"
	"klyvi-api/internal/movies"
	"klyvi-api/internal/platform/http/middleware"
	"klyvi-api/internal/search"
	"klyvi-api/internal/tv"
)

type Services struct {
	Movies *movies.Service
	TV     *tv.Service
	Search *search.Service
}

func New(services Services) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chimiddleware.RequestID)
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.StatsMiddleware)

	r.Get("/health", health.Get)

	r.Route("/v1", func(r chi.Router) {
		movieAPI := movies.NewAPI(services.Movies)
		r.Route("/movies", func(r chi.Router) {
			r.Get("/{id}", movieAPI.GetMovieById)
			r.Get("/{id}/recommendations", movieAPI.GetMovieRecommendations)
			r.Get("/{id}/collection", movieAPI.GetMovieCollection)
			r.Get("/", movieAPI.GetMovieList)
		})

		tvAPI := tv.NewAPI(services.TV)
		r.Route("/tv", func(r chi.Router) {
			r.Get("/{id}", tvAPI.GetTvById)
			r.Get("/{id}/recommendations", tvAPI.GetTvRecommendations)
			r.Get("/{id}/collection", tvAPI.GetTvCollection)
			r.Get("/", tvAPI.GetTvList)
		})

		searchAPI := search.NewAPI(services.Search)
		r.Get("/search", searchAPI.GetSearchResult)
	})

	return r
}

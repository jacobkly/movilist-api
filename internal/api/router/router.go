package router

import (
	"github.com/go-chi/chi/v5"

	"movilist-api/config"
	"movilist-api/internal/api/resource/health"
	"movilist-api/internal/api/resource/movies"
	"movilist-api/internal/api/resource/search"
	"movilist-api/internal/api/resource/tv"
	"movilist-api/pkg/tmdb"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Get)

	c := config.NewTMDB()
	tmdbClient := tmdb.NewClient(c.APIKey)

	r.Route("/v1", func(r chi.Router) {
		movieService := movies.NewService(tmdbClient)
		movieAPI := movies.NewAPI(movieService)

		r.Route("/movies", func(r chi.Router) {
			r.Get("/{id}", movieAPI.GetMovieById)
			r.Get("/{id}/recommendations", movieAPI.GetMovieRecommendations)
			r.Get("/{id}/collection", movieAPI.GetMovieCollection)
			r.Get("/", movieAPI.GetMovieList)
		})

		tvService := tv.NewService(tmdbClient)
		tvAPI := tv.NewAPI(tvService)

		r.Route("/tv", func(r chi.Router) {
			r.Get("/{id}", tvAPI.GetTvById)
			r.Get("/{id}/recommendations", tvAPI.GetTvRecommendations)
			r.Get("/{id}/collection", tvAPI.GetTvCollection)
			r.Get("/", tvAPI.GetTvList)
		})

		searchService := search.NewService(tmdbClient)
		searchAPI := search.NewAPI(searchService)

		r.Get("/search", searchAPI.GetSearchResult)
	})

	return r
}

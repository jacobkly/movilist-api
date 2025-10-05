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
			r.Get("/trending", movieAPI.GetTrendingMovies)
			r.Get("/upcoming", movieAPI.GetUpcomingMovies)
			r.Get("/popular", movieAPI.GetPopularMovies)
			r.Get("/top-rated", movieAPI.GetTopRatedMovies)
		})

		tvService := tv.NewService(tmdbClient)
		tvAPI := tv.NewAPI(tvService)

		r.Route("/tv", func(r chi.Router) {
			r.Get("/{id}", tvAPI.GetTvById)
			r.Get("/{id}/recommendations", tvAPI.GetTvRecommendations)
			r.Get("/trending", tvAPI.GetTrendingTv)
			r.Get("/upcoming", tvAPI.GetUpcomingTv)
			r.Get("/popular", tvAPI.GetPopularTv)
			r.Get("/top-rated", tvAPI.GetTopRatedTv)
		})

		searchService := search.NewService(tmdbClient)
		searchAPI := search.NewAPI(searchService)

		r.Get("/search", searchAPI.GetSearchResult)
	})

	return r
}

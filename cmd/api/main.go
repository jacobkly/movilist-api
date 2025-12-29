package main

import (
	"fmt"
	"log"
	"net/http"

	"movilist-api/config"
	"movilist-api/internal/api/resource/movies"
	"movilist-api/internal/api/resource/search"
	"movilist-api/internal/api/resource/tv"
	"movilist-api/internal/api/router"
	"movilist-api/internal/db"
	"movilist-api/pkg/tmdb"
)

//  @title          MoviList API
//  @version        1.0
//  @description    A RESTful API for managing movies, shows, and user media lists.
//                  Inspired by AniList, but focused on film and TV.

//  @contact.name   Jacob Klymenko
//  @contact.url    https://github.com/jacobkly

//  @license.name   MIT License
//  @license.url    https://opensource.org/licenses/MIT

// @host           localhost:8080
// @BasePath       /v1
func main() {
	cfg := config.New()

	dbConn, err := db.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	tmdbClient := tmdb.NewClient(cfg.TMDB.APIKey)

	movieRepo := db.NewMovieRepository(dbConn)

	movieService := movies.NewService(tmdbClient, movieRepo)
	tvService := tv.NewService(tmdbClient)
	searchService := search.NewService(tmdbClient)

	r := router.New(router.Services{
		Movies: movieService,
		TV:     tvService,
		Search: searchService,
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.TimeoutRead,
		WriteTimeout: cfg.Server.TimeoutWrite,
		IdleTimeout:  cfg.Server.TimeoutIdle,
	}

	log.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server startup failed:", err)
	}
}

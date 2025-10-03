package movies

import (
	"fmt"
	"movilist-api/pkg/tmdb"
)

type Service struct {
	client *tmdb.Client
}

func NewService(client *tmdb.Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetMovieById(id int) (interface{}, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?language=en-US", id)
	tmdbMovie, err := s.client.TMDBRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return tmdbMovie, nil
}

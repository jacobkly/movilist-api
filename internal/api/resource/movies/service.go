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
	endpoint := fmt.Sprintf("/movie/%d?language=en-US", id)
	movie, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Service) GetMovieRecommendations(id int) (interface{}, error) {
	endpoint := fmt.Sprintf("/movie/%d/recommendations?language=en-US&page=1", id)
	recommendations, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func (s *Service) GetMovieCollection(id int) (interface{}, error) {
	endpoint := fmt.Sprintf("/movie/%d?language=en-US", id)
	movie, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	belongsTo, ok := movie["belongs_to_collection"].(map[string]interface{})
	if !ok || belongsTo == nil {
		return nil, nil
	}

	collectionID, ok := belongsTo["id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid collection ID format")
	}

	endpoint = fmt.Sprintf("/collection/%d?language=en-US", int(collectionID))
	collection, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

// by weekly status
func (s *Service) GetTrendingMovies() (interface{}, error) {
	endpoint := "/trending/movie/week?language=en-US"
	trending, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return trending, nil
}

func (s *Service) GetUpcomingMovies() (interface{}, error) {
	endpoint := "/movie/upcoming?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

// by popularity
func (s *Service) GetPopularMovies() (interface{}, error) {
	endpoint := "/movie/popular?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

// by average score
func (s *Service) GetTopRatedMovies() (interface{}, error) {
	endpoint := "/movie/top_rated?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

func (s *Service) GetMovieList(listType string) (interface{}, error) {
	var endpoint string

	switch listType {
	case "trending":
		endpoint = "/trending/movie/week?language=en-US" // by weekly status
	case "upcoming":
		endpoint = "/movie/upcoming?language=en-US&page=1"
	case "popular":
		endpoint = "/movie/popular?language=en-US&page=1"
	case "top_rated":
		endpoint = "/movie/top_rated?language=en-US&page=1"
	default:
		return nil, fmt.Errorf("invalid list type")
	}

	return s.client.TMDBRequest("GET", endpoint, nil)
}

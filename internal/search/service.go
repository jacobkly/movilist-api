package search

import (
	"fmt"
	"net/url"
)

// TMDBClient is the upstream boundary for search. Search does not care about
// the concrete transport implementation; it only needs a component that can
// execute TMDB requests.
type TMDBClient interface {
	TMDBRequest(method, endpoint string, body interface{}) (map[string]interface{}, error)
}

type Service struct {
	client TMDBClient
}

// NewService wires search to the TMDB boundary through a small interface so
// the service can be tested without real network calls.
func NewService(client TMDBClient) *Service {
	return &Service{client: client}
}

func (s *Service) GetSearchResult(searchType string, query string) (interface{}, error) {
	var endpoint string

	switch searchType {
	case "movie":
		endpoint = "/search/movie"
	case "tv":
		endpoint = "/search/tv"
	case "person":
		endpoint = "/search/person"
	default:
		endpoint = "/search/multi"
	}

	endpoint = fmt.Sprintf("%s?query=%s&page=1&language=en-US", endpoint, url.QueryEscape(query))
	searchResult, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

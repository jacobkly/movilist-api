package search

import (
	"fmt"
	"movilist-api/pkg/tmdb"
	"net/url"
)

type Service struct {
	client *tmdb.Client
}

func NewService(client *tmdb.Client) *Service {
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

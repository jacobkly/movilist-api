package tv

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

// 0 fetches the entire tv show data, while others get exact season data
func (s *Service) GetTvById(idType string, id int, seasonNum int) (interface{}, error) {
	if idType == "internal" {
		return "internal ids have not yet been developed", nil
	}

	// external
	endpoint := fmt.Sprintf("/tv/%d?language=en-US", id)
	if seasonNum > 0 {
		endpoint = fmt.Sprintf("/tv/%d/season/%d?language=en-US", id, seasonNum)
	}

	tv, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return tv, nil
}

func (s *Service) GetTvRecommendations(idType string, id int) (interface{}, error) {
	if idType == "internal" {
		return "internal ids have not yet been developed", nil
	}

	endpoint := fmt.Sprintf("/tv/%d/recommendations?language=en-US&page=1", id)
	recommendations, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func (s *Service) GetTvCollection(idType string, id int) (interface{}, error) {
	if idType == "internal" {
		return "internal ids have not yet been developed", nil
	}

	endpoint := fmt.Sprintf("/tv/%d?language=en-US", id)
	tv, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	seasons, ok := tv["seasons"].([]interface{})
	if !ok || seasons == nil {
		return nil, nil
	}

	return seasons, nil
}

// by weekly status
func (s *Service) GetTrendingTv() (interface{}, error) {
	endpoint := "/trending/tv/week?language=en-US"
	trending, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return trending, nil
}

// tv that air in the next 7 days
func (s *Service) GetUpcomingTv() (interface{}, error) {
	endpoint := "/tv/on_the_air?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

// by popularity
func (s *Service) GetPopularTv() (interface{}, error) {
	endpoint := "/tv/popular?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

// by average score
func (s *Service) GetTopRatedTv() (interface{}, error) {
	endpoint := "/tv/top_rated?language=en-US&page=1"
	upcoming, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return upcoming, nil
}

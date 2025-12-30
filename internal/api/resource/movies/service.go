package movies

import (
	"context"
	"fmt"

	"movilist-api/internal/api/utils"
	"movilist-api/internal/db"
	"movilist-api/pkg/tmdb"
)

type Service struct {
	client *tmdb.Client
	repo   *db.MovieRepository
}

func NewService(client *tmdb.Client, repo *db.MovieRepository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func (s *Service) GetMovieById(ctx context.Context, id int, idType string) (*db.Movie, error) {
	var movieID int

	switch idType {
	case "tmdb":
		movieID = id

	case "media":
		var err error
		movieID, err = s.repo.GetMovieIDByMediaID(ctx, id)
		if err != nil {
			return nil, err
		}
		if movieID == 0 {
			return nil, fmt.Errorf("media not found")
		}

	default:
		return nil, fmt.Errorf("invalid id_type")
	}

	movie, err := s.repo.GetByTMDBID(ctx, movieID)
	if err != nil {
		return nil, err
	}
	if movie != nil {
		return movie, nil
	}

	raw, err := s.client.TMDBRequest(
		"GET",
		fmt.Sprintf("/movie/%d?language=en-US", movieID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	normalized := utils.NormalizeTMDBMovie(raw)
	if err := s.repo.InsertMovie(ctx, normalized); err != nil {
		return nil, err
	}

	_ = s.repo.EnsureMediaIndex(ctx, movieID)

	return normalized, nil
}

func (s *Service) GetMovieRecommendations(id int) (interface{}, error) {
	endpoint := fmt.Sprintf("/movie/%d/recommendations?language=en-US&page=1", id)
	recommendations, err := s.client.TMDBRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}

func (s *Service) GetMovieCollection(
	ctx context.Context,
	movieID int,
) ([]db.MovieCollection, error) {
	collectionID, err := s.repo.GetCollectionIDByMovieID(ctx, movieID)
	if err != nil {
		return nil, err
	}

	if collectionID != 0 {
		return s.repo.GetCollectionByCollectionID(ctx, collectionID)
	}

	rawMovie, err := s.client.TMDBRequest(
		"GET",
		fmt.Sprintf("/movie/%d?language=en-US", movieID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	belongsTo, ok := rawMovie["belongs_to_collection"].(map[string]interface{})
	if !ok || belongsTo == nil {
		return nil, nil
	}

	collectionID = int(belongsTo["id"].(float64))

	rawCollection, err := s.client.TMDBRequest(
		"GET",
		fmt.Sprintf("/collection/%d?language=en-US", collectionID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	collection := utils.NormalizeTMDBMovieCollection(rawCollection, collectionID)

	_ = s.repo.InsertMovieCollectionBatch(ctx, collection)

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

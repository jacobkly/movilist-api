package movies

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type API struct {
	service *Service
}

func NewAPI(service *Service) *API {
	return &API{service: service}
}

func (a *API) GetMovieById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	movie, err := a.service.GetMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (a *API) GetMovieRecommendations(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	recommendations, err := a.service.GetMovieRecommendations(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// as a collection for movies, it uses the collection system by TMDb
func (a *API) GetMovieCollection(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid movie id", http.StatusBadRequest)
		return
	}

	collection, err := a.service.GetMovieCollection(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if collection == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"details": "No collection found.",
		})
		return
	}
	json.NewEncoder(w).Encode(collection)
}

func (a *API) GetTrendingMovies(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetTrendingMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetUpcomingMovies(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetUpcomingMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetPopularMovies(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetPopularMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetTopRatedMovies(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetTopRatedMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetMovieList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	listType := query.Get("type")

	data, err := a.service.GetMovieList(listType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

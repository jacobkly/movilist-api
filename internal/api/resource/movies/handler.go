package movies

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

/*

GET    /v1/movies/trending           # trending movies
GET    /v1/movies/upcoming           # upcoming movies
GET    /v1/movies/popular            # popular movies (all time -> popularity, top 100 -> average score)

GET    /v1/movies/{id}               # full movie details
GET    /v1/movies/{id}/recommendations  # similar/recommended movies (last on priority)

*/

type API struct {
	service *Service
}

func NewAPI(service *Service) *API {
	return &API{service: service}
}

func (a *API) GetMovie(w http.ResponseWriter, r *http.Request) {
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

package tv

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

/*

for tv, it will all be with seasons (matching anilist behavior). to represent entire shows,
we do the first season.

GET    /v1/tv/trending           # trending tv shows (seasons)
GET    /v1/tv/upcoming           # upcoming tv shows (new seasons)
GET    /v1/tv/popular            # popular tv shows (all time -> popularity, top 100 -> average score)
													(first season shown as base)

# Series-level
GET    /v1/tv/{id}/recommendations  # similar/recommended series (last on priority)
	- have to figure something out as i will have first season represent an entire show. it's fine
	since every season in tmdb is just connected with series id and then season number so easy to
	find for any. the recommendation will just be generalized for the entire show not particular season.

# Season-level (for tracking/watchlist/ranking)
GET    /v1/tv/{id}     # full season details

*/

type API struct {
	service *Service
}

func NewAPI(service *Service) *API {
	return &API{service: service}
}

func (a *API) GetTvById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid tv id", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	idType := query.Get("id_type")
	seasonNum := 0

	/*
		will have to handle two api calls (if db hit is a miss) and then insert into db some new rows

		this is becoming more complex as we don't want to store same overall tv show data into potentially
		5 different media table rows (if tv show has more than one season -- highly likely)

		potential solution thought of right now is having the media table contain data on the season
		for the tv show, and then having another table for tv show information. when returning the response
		for the request, we will have to join the two queries. this avoids duplicate content amongst rows
		and serves as a good caching layer in the future. (even better is redis)
	*/

	if idType == "external" {
		// for now it skips if no season num given if external, but check api and see what happens
		// and see how it deals with one season total tv series. after checking, it's complex and view
		// comment above about how i would deal with it in db layer
		if s := query.Get("season_num"); s != "" {
			seasonNum, err = strconv.Atoi(s)
			if err != nil {
				http.Error(w, "invalid season number", http.StatusBadRequest)
				return
			}
		}
	}

	tv, err := a.service.GetTvById(idType, id, seasonNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tv)
}

func (a *API) GetTvRecommendations(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid tv id", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	idType := query.Get("id_type")

	recommendations, err := a.service.GetTvRecommendations(idType, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

func (a *API) GetTrendingTv(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetTrendingTv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetUpcomingTv(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetUpcomingTv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetPopularTv(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetPopularTv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

func (a *API) GetTopRatedTv(w http.ResponseWriter, r *http.Request) {
	trending, err := a.service.GetTopRatedTv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trending)
}

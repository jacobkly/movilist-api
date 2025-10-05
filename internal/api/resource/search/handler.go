package search

import (
	"encoding/json"
	"net/http"
)

type API struct {
	service *Service
}

func NewAPI(service *Service) *API {
	return &API{service: service}
}

func (a *API) GetSearchResult(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	searchType := queryParams.Get("type")
	// TODO: can do enum check for the search type
	query := queryParams.Get("query")

	searchResult, err := a.service.GetSearchResult(searchType, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResult)
}

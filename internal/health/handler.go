package health

import (
	"net/http"

	"klyvi-api/internal/platform/http/middleware"
	"klyvi-api/internal/platform/http/response"
)

func Get(w http.ResponseWriter, r *http.Request) {
	response.WriteSuccess(
		w,
		http.StatusOK,
		"v1",
		middleware.StatsFromContext(r.Context()),
		map[string]interface{}{
			"status_code": 200,
			"details":     "Server is up and running.",
		},
	)
}

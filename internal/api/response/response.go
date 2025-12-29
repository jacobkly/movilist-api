package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Version string      `json:"version"`
	Stats   interface{} `json:"stats,omitempty"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteSuccess(
	w http.ResponseWriter,
	status int,
	version string,
	stats interface{},
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(SuccessResponse{
		Version: version,
		Stats:   stats,
		Data:    data,
	})
}

func WriteError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}

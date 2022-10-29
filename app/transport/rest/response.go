package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response will wrap message
// that will be sent in JSON format.
type Response struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// WriteJSONResponse writes JSON response.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("failed encoding to JSON with status code : ")
	}
}

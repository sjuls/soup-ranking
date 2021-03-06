package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON is a simple helper function to send a JSON response object
// through a response writer.
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

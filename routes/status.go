package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Status struct {
	Status      string    `json:"status"`
}

func AddStatus(r *mux.Router) {
	r.Methods("GET").
		Name("Status").
		Path("/_status").
		HandlerFunc(statusHandlerFunc)
}

func statusHandlerFunc(w http.ResponseWriter, r *http.Request) {
	status := Status{
		"I am ALIVE!",
	}
	JSON(w, status)
}
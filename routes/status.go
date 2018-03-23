package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Status struct {
	Status string `json:"status"`
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

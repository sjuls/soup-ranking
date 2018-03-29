package status

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/utils"
)

type status struct {
	Status string `json:"status"`
}

// AddRoute - adds a status route to the provided router
func AddRoute(r *mux.Router) {
	r.Methods("GET").
		Name("Status").
		Path("/_status").
		HandlerFunc(statusHandlerFunc)
}

func statusHandlerFunc(w http.ResponseWriter, r *http.Request) {
	status := status{
		"I am ALIVE!",
	}

	utils.JSON(w, status)
}

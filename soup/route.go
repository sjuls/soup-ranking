package soup

import (
	"net/http"

	"github.com/sjuls/soup-ranking/utils"

	"github.com/gorilla/mux"
)

const (
	path = "/soup"
)

type soupDto struct {
	Name *string `json:"name"`
}

// AddRoute - Adds routes to the provided router
// to enable the fetching of soups
func AddRoute(soupManager *Manager) func(r *mux.Router) {
	return func(r *mux.Router) {
		r.Methods("GET").
			Name("Get the soup of the day").
			Path(path).
			HandlerFunc(createGetSoupHandlerFunc(soupManager))
	}
}

func createGetSoupHandlerFunc(soupManager *Manager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		soupName, err := soupManager.GetSoupName()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		dto := soupDto{Name: soupName}

		utils.WriteJSON(w, dto)
	}
}

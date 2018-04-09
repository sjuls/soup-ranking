package soup

import (
	"net/http"

	"github.com/sjuls/soup-ranking/utils"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
)

const (
	path = "/soup"
)

type soupDto struct {
	Name string `json:"name"`
}

// AddRoute - Adds routes to the provided router
// to enable the fetching of soups
func AddRoute(repo dbctx.SoupRepository) func(r *mux.Router) {
	return func(r *mux.Router) {
		r.Methods("GET").
			Name("Get the soup of the day").
			Path(path).
			HandlerFunc(createGetSoupHandlerFunc(repo))
	}
}

func createGetSoupHandlerFunc(repo dbctx.SoupRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		soupOfTheDay, err := repo.GetSoupOfTheDay()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		dto := soupDto{Name: soupOfTheDay.Soup.Name}

		utils.WriteJSON(w, dto)
	}
}

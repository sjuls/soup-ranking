package score

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
	"github.com/sjuls/soup-ranking/utils"
)

const (
	path = "/score"
)

type scoreDto struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// AddRoute - Adds routes to the provided router
// to enable addition and fetching of soup scores
func AddRoute(r *mux.Router) {
	r.Methods("GET").
		Name("Get score").
		Path(path).
		HandlerFunc(getScoreHandlerFunc)

	r.Methods("POST").
		Name("Add score").
		Path(path).
		HandlerFunc(postScoreHandlerFunc)
}

func getScoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	db, err := dbctx.Open()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var scores []dbctx.Score

	db.Find(&scores)

	utils.JSON(w, scores)
}

func postScoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	score := &scoreDto{}
	if err := json.NewDecoder(r.Body).Decode(score); err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	db, err := dbctx.Open()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	db.Create(&dbctx.Score{
		Score:   score.Score,
		Comment: score.Comment,
	})

	w.WriteHeader(http.StatusAccepted)
}

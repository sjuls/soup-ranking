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

var (
	repo Repository
)

type scoreDto struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// AddRoute - Adds routes to the provided router
// to enable addition and fetching of soup scores
func AddRoute(r *mux.Router) {
	repo = Repository{}
	r.Methods("GET").
		Name("Get score").
		Path(path).
		HandlerFunc(getScoresHandlerFunc)

	r.Methods("POST").
		Name("Add score").
		Path(path).
		HandlerFunc(postScoreHandlerFunc)
}

func getScoresHandlerFunc(w http.ResponseWriter, r *http.Request) {
	scores, err := repo.GetScores()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.JSON(w, scores)
}

func postScoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	score := &scoreDto{}
	if err := json.NewDecoder(r.Body).Decode(score); err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	err := repo.SaveScore(&dbctx.Score{
		Score:   score.Score,
		Comment: score.Comment,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

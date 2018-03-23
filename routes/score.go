package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
	"net/http"
)

const (
	path = "/score"
)

type ScoreDto struct {
	Score int `json:"score"`
}

func AddScore(r *mux.Router) {
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

	JSON(w, scores)
}

func postScoreHandlerFunc(w http.ResponseWriter, r *http.Request) {
	db, err := dbctx.Open()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	score := &ScoreDto{}
	err = json.NewDecoder(r.Body).Decode(score)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	db.Create(&dbctx.Score{
		Score: score.Score,
	})
	w.WriteHeader(http.StatusAccepted)
}

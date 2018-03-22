package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
)

const (
	path = "/score"
)

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

	score := &dbctx.Score{}
	err = json.NewDecoder(r.Body).Decode(score)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	
	db.Create(score)
	w.WriteHeader(http.StatusAccepted)
}
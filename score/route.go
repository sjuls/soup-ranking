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
	Score   int     `json:"score"`
	Comment *string `json:"comment"`
}

// AddRoute - Adds routes to the provided router
// to enable addition and fetching of soup scores
func AddRoute(repo dbctx.ScoreRepository) func(r *mux.Router) {
	return func(r *mux.Router) {
		r.Methods("GET").
			Name("Get score").
			Path(path).
			HandlerFunc(createGetScoresHandlerFunc(repo))

		r.Methods("POST").
			Name("Add score").
			Path(path).
			HandlerFunc(createPostScoreHandlerFunc(repo))
	}
}

func createGetScoresHandlerFunc(repo dbctx.ScoreRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		scores, err := repo.GetScores()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		dtos := make([]scoreDto, len(*scores))
		for i, score := range *scores {
			dtos[i] = scoreDto{
				Score:   score.Score,
				Comment: score.Comment,
			}
		}

		utils.WriteJSON(w, &dtos)
	}
}

func createPostScoreHandlerFunc(repo dbctx.ScoreRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

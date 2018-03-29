package slack

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/utils"
)

const (
	challengeType string = "url_verification"
)

var (
	appToken string
)

type (
	challengeResponse struct {
		Challenge string `json:"challenge"`
	}
)

// AddRoute - adds a slack event route which accepts event with the given token
func AddRoute(token string) func(r *mux.Router) {
	appToken = token
	return func(r *mux.Router) {
		r.Methods("POST").
			Name("Slack").
			Path("/slack").
			HandlerFunc(slackHandlerFunc)
	}
}

func slackHandlerFunc(w http.ResponseWriter, r *http.Request) {
	event := &Event{}
	if err := json.NewDecoder(r.Body).Decode(event); err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	if event.Token != appToken {
		log.Printf("Invalid Slack token '%s'!", event.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if event.Type == challengeType {
		log.Println("Slack challenge received!")
		utils.JSON(w, challengeResponse{event.Challenge})
		return
	}

	go HandleEvent(event)

	w.WriteHeader(http.StatusOK)
}

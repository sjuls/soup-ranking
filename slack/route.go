package slack

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sjuls/soup-ranking/score"
	"github.com/sjuls/soup-ranking/soup"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/slack/api"
	"github.com/sjuls/soup-ranking/utils"
)

const (
	challengeType string = "url_verification"
)

type (
	challengeResponse struct {
		Challenge string `json:"challenge"`
	}
)

// AddRoute - adds a slack event route which accepts event with the given token
func AddRoute(verificationToken string, baseURL string, accessToken string, soupRepository soup.Repository, scoreRepository score.Repository, adminUsers []string) func(r *mux.Router) {
	webAPI := api.NewSlackWebAPI(baseURL, accessToken, &http.Client{})
	globalEventHandler := &GlobalEventHandler{
		[]EventHandler{
			NewCommandsHandler(webAPI, soupRepository, scoreRepository, adminUsers),
		},
	}

	return func(r *mux.Router) {
		r.Methods("POST").
			Name("Slack").
			Path("/slack/event").
			HandlerFunc(createSlackHandlerFunc(verificationToken, globalEventHandler))
	}
}

func createSlackHandlerFunc(verificationToken string, eventHandler EventHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		event := &EventCallback{}
		if err := json.NewDecoder(r.Body).Decode(event); err != nil {
			log.Printf("SlackHandler error %s", err)
			http.Error(w, "Cannot read body", http.StatusBadRequest)
			return
		}

		if event.Token != verificationToken {
			log.Printf("Invalid Slack token '%s'!", event.Token)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if event.Type == challengeType {
			log.Println("Slack challenge received!")
			utils.JSON(w, challengeResponse{event.Challenge})
			return
		}

		go eventHandler.HandleEvent(event)

		w.WriteHeader(http.StatusOK)
	}
}

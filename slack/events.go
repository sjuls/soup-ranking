package slack

import (
	"log"
)

// Event - represents an event sent by Slack
type Event struct {
	Token       string              `json:"token"`
	Challenge   string              `json:"challenge"`
	TeamID      string              `json:"team_id"`
	APIAppID    string              `json:"api_app_id"`
	Event       map[string]struct{} `json:"event"`
	Type        string              `json:"type"`
	AuthedUsers []string            `json:"authed_users"`
	EventID     string              `json:"event_id"`
	EventTime   int                 `json:"event_time"`
}

// HandleEvent - Handle a Slack event.
func HandleEvent(event *Event) {
	// TODO: Handle events...
	log.Println("Received Slack event")
}

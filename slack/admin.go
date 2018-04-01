package slack

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/sjuls/soup-ranking/slack/api"
)

type (
	// AdminHandler handles admin commands received through Slack events
	AdminHandler struct {
		WebAPI     *api.SlackWebAPI
		AdminUsers []string
	}
)

// HandleEvent handles events delegated to AdminHandler
func (h *AdminHandler) HandleEvent(event *EventCallback) {
	if !shouldHandle(event) {
		return
	}

	innerEvent := DMEvent{}
	if err := mapstructure.Decode(event.Event, &innerEvent); err != nil {
		panic(err)
	}

	// TODO: Do something with the events
	message := api.PostMessage{
		Channel: innerEvent.Channel,
		Text:    fmt.Sprintf("Echo: %s", innerEvent.Text),
		AsUser:  true,
	}

	_, err := h.WebAPI.PostMessage(&message)
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func shouldHandle(event *EventCallback) bool {
	if event.Event["type"] != MessageEventType {
		return false
	}

	// Do not handle events triggered by the bots themselves.
	for _, authedUser := range event.AuthedUsers {
		if event.Event["user"] == authedUser {
			return false
		}
	}

	return true
}

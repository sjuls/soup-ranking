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
func (h *AdminHandler) HandleEvent(event *Event) {
	if event.Type != DMEventType {
		return
	}

	innerEvent := DMEvent{}
	if err := mapstructure.Decode(event.Event, &innerEvent); err != nil {
		log.Print(err)
		panic(err)
	}

	message := api.PostMessage{
		Channel: innerEvent.Channel,
		Text:    fmt.Sprintf("echo: %s", innerEvent.Text),
	}

	if _, err := h.WebAPI.PostMessage(&message); err != nil {
		log.Print(err)
	}
}

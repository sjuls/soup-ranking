package slack

type (
	// GlobalEventHandler propogates calls to registered eventhandlers
	GlobalEventHandler struct {
		EventHandlers []EventHandler
	}

	// EventHandler interface is used to register event handlers for Slack events.
	EventHandler interface {
		HandleEvent(event *Event)
	}
)

// HandleEvent - Handle a Slack event by triggering registered eventhandlers.
func (geh *GlobalEventHandler) HandleEvent(event *Event) {
	for _, eventHandler := range geh.EventHandlers {
		go eventHandler.HandleEvent(event)
	}
}

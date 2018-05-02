package slack

type (
	// GlobalEventHandler propogates calls to registered eventhandlers
	GlobalEventHandler struct {
		EventHandlers []EventHandler
	}

	// EventHandler interface is used to register event handlers for Slack events.
	EventHandler interface {
		HandleEvent(event *EventCallback)
	}
)

var (
	handledEvents = map[string]int{}
)

// HandleEvent - Handle a Slack event by triggering registered eventhandlers.
func (geh *GlobalEventHandler) HandleEvent(event *EventCallback) {
	// Only handle events not seen before. TODO: Make scalable.
	if _, ok := handledEvents[event.EventID]; !ok {
		handledEvents[event.EventID] = event.EventTime
		for _, eventHandler := range geh.EventHandlers {
			go eventHandler.HandleEvent(event)
		}
	}
}

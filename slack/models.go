package slack

const (
	// MessageEventType holds the direct message event type
	MessageEventType = "message"
)

type (
	// EventCallback - the outer event wrapper that represents an event sent by Slack
	EventCallback struct {
		Token       string                 `json:"token"`
		Challenge   string                 `json:"challenge"`
		TeamID      string                 `json:"team_id"`
		APIAppID    string                 `json:"api_app_id"`
		Event       map[string]interface{} `json:"event"`
		Type        string                 `json:"type"`
		AuthedUsers []string               `json:"authed_users"`
		EventID     string                 `json:"event_id"`
		EventTime   int                    `json:"event_time"`
	}

	// MessageEvent - the inner event of a Slack message event callback
	MessageEvent struct {
		Type      string `json:"type"`
		Channel   string `json:"channel"`
		User      string `json:"user"`
		Text      string `json:"text"`
		TimeStamp string `json:"ts"`
	}
)

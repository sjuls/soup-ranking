package slack

const (
	// DMEventType holds the direct message event type
	DMEventType = "message.im"
)

type (
	// Event - the outer event wrapper that represents an event sent by Slack
	Event struct {
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

	// DMEvent - the inner event of a Slack direct message event (message.im)
	DMEvent struct {
		Type      string `json:"type"`
		Channel   string `json:"channel"`
		User      string `json:"user"`
		Text      string `json:"text"`
		TimeStamp string `json:"ts"`
	}
)

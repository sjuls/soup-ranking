package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// SlackWebAPI simplifies interactions with Slack's web api, see <a href="https://api.slack.com/web" />
	SlackWebAPI struct {
		BaseURL     string
		AccessToken string
	}

	// PostMessage struct is the data transfer object available for the chat.postMessage method
	PostMessage struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
		AsUser  bool   `json:"as_user"`
		// TODO: Optional properties
	}
)

// PostMessage invokes the chat.postMessage method of the Slack web api
func (api *SlackWebAPI) PostMessage(message *PostMessage) (*http.Response, error) {
	return api.send("chat.postMessage", message)
}

func (api *SlackWebAPI) send(action string, v interface{}) (*http.Response, error) {
	jsonValue, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", api.BaseURL, action), bytes.NewReader(jsonValue))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.AccessToken))
	request.Header.Set("Content-type", "application/json; charset=utf-8")

	client := &http.Client{}
	return client.Do(request)
}

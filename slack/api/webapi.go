package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sjuls/soup-ranking/utils"
)

type (
	slackWebAPI struct {
		baseURL     string
		accessToken string
		httpClient  utils.HTTPClient
	}

	// SlackWebAPI simplifies interactions with Slack's web api, see <a href="https://api.slack.com/web" />
	SlackWebAPI interface {
		PostMessage(message *PostMessage) (*http.Response, error)
	}

	// PostMessage struct is the data transfer object available for the chat.postMessage method
	PostMessage struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
		AsUser  bool   `json:"as_user"`
		// TODO: Optional properties
	}
)

// NewSlackWebAPI creates a new SlackWebApi
func NewSlackWebAPI(baseURL string, accessToken string, httpClient utils.HTTPClient) SlackWebAPI {
	return &slackWebAPI{
		baseURL,
		accessToken,
		httpClient,
	}
}

// PostMessage invokes the chat.postMessage method of the Slack web api
func (api *slackWebAPI) PostMessage(message *PostMessage) (*http.Response, error) {
	return api.send("chat.postMessage", message)
}

func (api *slackWebAPI) send(action string, v interface{}) (*http.Response, error) {
	jsonValue, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", api.baseURL, action), bytes.NewReader(jsonValue))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.accessToken))
	request.Header.Set("Content-type", "application/json; charset=utf-8")

	return api.httpClient.Do(request)
}

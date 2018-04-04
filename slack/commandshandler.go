package slack

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sjuls/soup-ranking/slack/api"
	"github.com/sjuls/soup-ranking/slack/commands"
)

type (
	// CommandsHandler handles admin commands received through Slack events
	CommandsHandler struct {
		WebAPI     *api.SlackWebAPI
		AdminUsers []string
	}
)

const (
	maxFlags = 25 // Maximum number of flags that will be extracted
)

var (
	registeredCommands = map[string]func(flags []string, output io.Writer){
		"soup": commands.Soup,
		"rate": commands.Rate,
	}
	adminCommands = []string{
		"soup",
	}
)

// HandleEvent handles events delegated to AdminHandler
func (h *CommandsHandler) HandleEvent(event *EventCallback) {
	if !shouldHandle(event) {
		return
	}

	innerEvent := MessageEvent{}
	if err := mapstructure.Decode(event.Event, &innerEvent); err != nil {
		panic(err)
	}

	cmdRegex, _ := regexp.Compile(fmt.Sprintf("^\\s*(?:<@(?:%s)>)\\s+(\\S+)\\s+(.*)$", strings.Join(event.AuthedUsers, "|")))
	cmdMatches := cmdRegex.FindStringSubmatch(innerEvent.Text)

	if len(cmdMatches) < 3 {
		return
	}

	flagRegex, _ := regexp.Compile("([^\\s\"]+)|\"([^\"]+)\"")
	flags := extractFlags(flagRegex.FindAllStringSubmatch(cmdMatches[2], maxFlags))

	commandName := strings.ToLower(cmdMatches[1])

	output := bytes.NewBuffer([]byte{})
	if h.isAuthorized(commandName, innerEvent.User) {
		h.executeCommand(commandName, flags, output)
	} else {
		fmt.Fprintf(output, "Sorry <@%s>, you're not authorized to invoke that command.", innerEvent.User)
	}

	message := api.PostMessage{
		Channel: innerEvent.Channel,
		Text:    output.String(),
		AsUser:  true,
	}

	if _, err := h.WebAPI.PostMessage(&message); err != nil {
		panic(err)
	}
}

func (h *CommandsHandler) isAuthorized(commandName string, user string) bool {
	for _, adminCommand := range adminCommands {
		if adminCommand == commandName {
			return h.isAdminUser(user)
		}
	}
	return true
}

func (h *CommandsHandler) isAdminUser(user string) bool {
	for _, adminUser := range h.AdminUsers {
		if adminUser == user {
			return true
		}
	}
	return false
}

func (h *CommandsHandler) executeCommand(commandName string, flags []string, output io.Writer) {
	if command := registeredCommands[commandName]; command != nil {
		command(flags, output)
	}
}

func extractFlags(flagMatches [][]string) []string {
	flags := make([]string, len(flagMatches))
	for i, flag := range flagMatches {
		flags[i] = flag[1]
	}
	return flags
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

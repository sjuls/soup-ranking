package slack

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/sjuls/soup-ranking/score"

	"github.com/mitchellh/mapstructure"
	"github.com/sjuls/soup-ranking/slack/api"
	"github.com/sjuls/soup-ranking/slack/commands"
)

type (
	// CommandsHandler handles admin commands received through Slack events
	CommandsHandler struct {
		webAPI             api.SlackWebAPI
		adminUsers         []string
		registeredCommands map[string]commands.Command
	}
)

const (
	maxFlags = 25 // Maximum number of flags that will be extracted
)

// NewCommandsHandler create a new commandshandler
func NewCommandsHandler(webAPI api.SlackWebAPI, repo score.Repository, adminUsers []string) EventHandler {
	commands := map[string]commands.Command{
		"today": commands.NewTodayCommand(),
		"rate":  commands.NewRateCommand(repo),
	}

	var handler EventHandler = &CommandsHandler{
		webAPI,
		adminUsers,
		commands,
	}

	return handler
}

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

	if _, err := h.webAPI.PostMessage(&message); err != nil {
		panic(err)
	}
}

func (h *CommandsHandler) isAuthorized(commandName string, user string) bool {
	if h.registeredCommands[commandName].RequiresAdmin() {
		return h.isAdminUser(user)
	}
	return true
}

func (h *CommandsHandler) isAdminUser(user string) bool {
	for _, adminUser := range h.adminUsers {
		if adminUser == user {
			return true
		}
	}
	return false
}

func (h *CommandsHandler) executeCommand(commandName string, flags []string, output io.Writer) {
	if command := h.registeredCommands[commandName]; command != nil {
		command.Execute(flags, output)
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

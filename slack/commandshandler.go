package slack

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/sjuls/soup-ranking/dbctx"

	"github.com/mitchellh/mapstructure"
	"github.com/sjuls/soup-ranking/slack/api"
	"github.com/sjuls/soup-ranking/slack/commands"
)

type (
	commandsHandler struct {
		webAPI             api.SlackWebAPI
		adminUsers         []string
		registeredCommands map[string]commands.Command
	}
)

const (
	maxFlags     = 25 // Maximum number of flags that will be extracted
	usageCommand = "usage"
)

// NewCommandsHandler create a new commandshandler
func NewCommandsHandler(
	webAPI api.SlackWebAPI,
	soupRepository dbctx.SoupRepository,
	scoreRepository dbctx.ScoreRepository,
	adminUsers []string,
) EventHandler {
	commands := map[string]commands.Command{
		"today":     commands.NewTodayCommand(soupRepository),
		"set-today": commands.NewSetTodayCommand(soupRepository),
		"rate":      commands.NewRateCommand(scoreRepository),
	}

	return &commandsHandler{
		webAPI,
		adminUsers,
		commands,
	}
}

// HandleEvent handles events delegated to AdminHandler
func (h *commandsHandler) HandleEvent(event *EventCallback) {
	if !shouldHandle(event) {
		return
	}

	output := bytes.NewBuffer([]byte{})
	innerEvent := MessageEvent{}
	if err := mapstructure.Decode(event.Event, &innerEvent); err != nil {
		panic(err)
	}

	cmdRegex, _ := regexp.Compile(fmt.Sprintf("^\\s*(?:<@(?:%s)>)\\s+(\\S+)\\s?(.*)$", strings.Join(event.AuthedUsers, "|")))
	cmdMatches := cmdRegex.FindStringSubmatch(innerEvent.Text)

	if len(cmdMatches) < 2 {
		return
	}

	commandName := strings.ToLower(cmdMatches[1])

	var args string
	if len(cmdMatches) == 3 {
		args = strings.TrimSpace(cmdMatches[2])
	}

	if commandName == usageCommand {
		fmt.Fprintf(output, "Hello <@%s>, here's a list of the commands available to you - enjoy!\n\n", innerEvent.User)
		h.Usage(output, innerEvent.User)
	} else if h.isAuthorized(commandName, innerEvent.User) {
		h.executeCommand(commandName, args, output)
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

func (h *commandsHandler) Usage(output io.Writer, user string) {
	for _, command := range h.registeredCommands {
		if !command.RequiresAdmin() || h.isAdminUser(user) {
			fmt.Fprintln(output, command.Usage())
		}
	}
}

func (h *commandsHandler) isAuthorized(commandName string, user string) bool {
	if command, ok := h.registeredCommands[commandName]; ok && command.RequiresAdmin() {
		return h.isAdminUser(user)
	}
	return true
}

func (h *commandsHandler) isAdminUser(user string) bool {
	for _, adminUser := range h.adminUsers {
		if adminUser == user {
			return true
		}
	}
	return false
}

func (h *commandsHandler) executeCommand(commandName string, args string, output io.Writer) {
	if command, ok := h.registeredCommands[commandName]; ok {
		command.Execute(args, output)
	} else {
		fmt.Fprintf(output, "Sorry, I didn't recognize command '%s'. Please ask for 'usage' to see the available commands.", commandName)
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

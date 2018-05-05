package commands

import (
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/soup"
)

type (
	syncSoupCommand struct {
		soupManager *soup.Manager
	}
)

// NewSyncSoupCommand create a new sync-soup command
func NewSyncSoupCommand(soupManager *soup.Manager) Command {
	return &syncSoupCommand{
		soupManager,
	}
}

func (c *syncSoupCommand) Execute(args string, output io.Writer) {
	soupName, err := c.soupManager.SyncSoupName()
	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintf(output, "Roger, admin. The soup of the day is now synced with the menu as: `%s`.", *soupName)
	}
}

func (c *syncSoupCommand) RequiresAdmin() bool {
	return true
}

func (c *syncSoupCommand) Usage() string {
	return "`<@soupbot> sync-soup` Force sync the soup of the day with the menu."
}

package commands

import (
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/soup"
)

type (
	todayCommand struct {
		soupManager *soup.Manager
	}
)

// NewTodayCommand create a new today command
func NewTodayCommand(soupManager *soup.Manager) Command {
	return &todayCommand{
		soupManager,
	}
}

func (c *todayCommand) Execute(args string, output io.Writer) {
	soupName, err := c.soupManager.GetSoupName()
	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintf(output, "The soup of the day is: `%s`.", *soupName)
	}
}

func (c *todayCommand) RequiresAdmin() bool {
	return false
}

func (c *todayCommand) Usage() string {
	return "`<@soupbot> today` Returns the soup of the day"
}

package commands

import (
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	todayCommand struct {
		repo dbctx.SoupRepository
	}
)

// NewTodayCommand create a new today command
func NewTodayCommand(repo dbctx.SoupRepository) Command {
	return &todayCommand{
		repo,
	}
}

func (c *todayCommand) Execute(args string, output io.Writer) {
	soupOfTheDay, err := c.repo.GetSoupOfTheDay()
	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintf(output, "The soup of the day is: %s.", soupOfTheDay.Soup.Name)
	}
}

func (c *todayCommand) RequiresAdmin() bool {
	return false
}

func (c *todayCommand) Usage() string {
	return "`<@soupbot> today` Returns the soup of the day"
}

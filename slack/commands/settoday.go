package commands

import (
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	setTodayCommand struct {
		repo dbctx.SoupRepository
	}
)

// NewSetTodayCommand create a new set-today command
func NewSetTodayCommand(repo dbctx.SoupRepository) Command {
	return &setTodayCommand{
		repo,
	}
}

func (c *setTodayCommand) Execute(args string, output io.Writer) {
	err := c.repo.SetSoup(args)
	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintf(output, "Roger, admin. The soup of the day is now set to: %s.", args)
	}
}

func (c *setTodayCommand) RequiresAdmin() bool {
	return true
}

func (c *setTodayCommand) Usage() string {
	return "`<@soupbot> set-today <name of the soup of the day>` Set the name of the soup of the day."
}

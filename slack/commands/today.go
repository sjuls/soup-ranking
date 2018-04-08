package commands

import (
	"flag"
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/soup"
	"github.com/sjuls/soup-ranking/utils"
)

type (
	todayFlags struct {
		Name string
	}

	todayCommand struct {
		repo soup.Repository
	}
)

// NewTodayCommand create a new today command
func NewTodayCommand(repo soup.Repository) Command {
	return &todayCommand{
		repo,
	}
}

func (c *todayCommand) Execute(args []string, output io.Writer) {
	flags, err := extractTodayFlags(args, output)
	if err != nil {
		return
	}

	err = c.repo.SetSoup(flags.Name)

	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintf(output, "Roger. Soup of the day is now set to '%s'.", flags.Name)
	}
}

func (c *todayCommand) RequiresAdmin() bool {
	return true
}

func extractTodayFlags(args []string, output io.Writer) (*todayFlags, error) {
	flags := todayFlags{}
	config := func(flagset *flag.FlagSet) {
		flagset.StringVar(&flags.Name, "name", "Unknown soup", "Set the name of the soup of the day")
	}

	if err := utils.ParseArguments("today", args, config, output); err != nil {
		return nil, err
	}

	return &flags, nil
}

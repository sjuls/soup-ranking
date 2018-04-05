package commands

import (
	"flag"
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/utils"
)

type (
	todayFlags struct {
		Name string
	}

	todayCommand struct {
	}
)

// NewTodayCommand create a new today command
func NewTodayCommand() Command {
	var todayCmd Command = &todayCommand{}
	return todayCmd
}

func (c *todayCommand) Execute(args []string, output io.Writer) {
	flags, err := extractTodayFlags(args, output)
	if err != nil {
		return
	}

	fmt.Fprintln(output, flags.Name)
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

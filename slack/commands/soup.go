package commands

import (
	"flag"
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/utils"
)

type (
	soupFlags struct {
		Name string
	}
)

// Soup command function for the soup command
func Soup(args []string, output io.Writer) {
	flags, err := extractSoupFlags(args, output)
	if err != nil {
		return
	}

	fmt.Fprintln(output, flags.Name)
}

func extractSoupFlags(args []string, output io.Writer) (*soupFlags, error) {
	flags := soupFlags{}
	config := func(flagset *flag.FlagSet) {
		flagset.StringVar(&flags.Name, "name", "Unknown soup", "Set the name of the soup of the day")
	}

	if err := utils.ParseArguments("soup", args, config, output); err != nil {
		return nil, err
	}

	return &flags, nil
}

package commands

import (
	"flag"
	"fmt"
	"io"

	"github.com/sjuls/soup-ranking/utils"
)

type (
	soupFlags struct {
		Title string
	}
)

// Soup command function for the soup command
func Soup(args []string, output io.Writer) {
	flags, err := extractFlags(args, output)
	if err != nil {
		return
	}

	fmt.Fprintln(output, flags.Title)
}

func extractFlags(args []string, output io.Writer) (*soupFlags, error) {
	flags := soupFlags{}
	config := func(flagset *flag.FlagSet) {
		flagset.StringVar(&flags.Title, "title", "Unknown soup", "Set the title of the soup of the day")
	}

	if err := utils.ParseArguments("soup", args, config, output); err != nil {
		return nil, err
	}

	return &flags, nil
}

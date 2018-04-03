package utils

import (
	"flag"
	"fmt"
	"io"
)

// ParseArguments parses the given arguments with the configuration function and pipes output to the writer
func ParseArguments(name string, args []string, configureFlags func(*flag.FlagSet), output io.Writer) error {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.SetOutput(output)

	configureFlags(flags)

	if err := flags.Parse(args); err != nil {
		fmt.Fprintln(output, err.Error())
		return err
	}

	return nil
}

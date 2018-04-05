package commands

import (
	"io"
)

type (
	// Command which can be executed
	Command interface {
		Execute(args []string, writer io.Writer)

		RequiresAdmin() bool
	}
)

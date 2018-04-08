package commands

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/sjuls/soup-ranking/dbctx"

	"github.com/sjuls/soup-ranking/score"
	"github.com/sjuls/soup-ranking/utils"
)

type (
	rateFlags struct {
		Score   int
		Comment string
	}

	rateCommand struct {
		repo score.Repository
	}
)

// NewRateCommand create a new rate command
func NewRateCommand(repo score.Repository) Command {
	return &rateCommand{repo}
}

func (c *rateCommand) Execute(args []string, output io.Writer) {
	flags, err := extractRateFlags(args, output)
	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	if 1 > flags.Score || flags.Score > 10 {
		fmt.Fprintln(output, "Score should be between 1 and 10")
		return
	}

	err = c.repo.SaveScore(&dbctx.Score{
		Score:   flags.Score,
		Comment: flags.Comment,
	})

	if err != nil {
		log.Println(err.Error())
		fmt.Fprintln(output, "An error has occurred.")
	} else {
		fmt.Fprintln(output, "Thank you for your soup rating!")
	}
}

func (c *rateCommand) RequiresAdmin() bool {
	return false
}

func extractRateFlags(args []string, output io.Writer) (*rateFlags, error) {
	flags := rateFlags{}
	config := func(flagset *flag.FlagSet) {
		flagset.IntVar(&flags.Score, "score", 0, "Choose a score from 1 to 10.")
		flagset.StringVar(&flags.Comment, "comment", "No comment", "Textual comment in case the score isn't enough for you.")
	}

	if err := utils.ParseArguments("rate", args, config, output); err != nil {
		return nil, err
	}

	return &flags, nil
}

package commands

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	rateFlags struct {
		Score   int
		Comment string
	}

	rateCommand struct {
		repo dbctx.ScoreRepository
	}
)

// NewRateCommand create a new rate command
func NewRateCommand(repo dbctx.ScoreRepository) Command {
	return &rateCommand{repo}
}

func (c *rateCommand) Execute(args string, output io.Writer) {
	rateMatcher, err := regexp.Compile("([1-9]|10)\\s(.*)")
	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	rateMatches := rateMatcher.FindStringSubmatch(args)
	scoreValue, err := strconv.ParseInt(rateMatches[1], 10, 32)

	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	if scoreValue < 1 || 10 < scoreValue {
		fmt.Fprintln(output, "Score must be between 1 and 10")
		return
	}

	score := dbctx.Score{
		Score: int(scoreValue),
	}

	if len(args) > 2 {
		score.Comment = rateMatches[2]
	}

	err = c.repo.SaveScore(&score)

	if err != nil {
		fmt.Fprintln(output, err.Error())
	} else {
		fmt.Fprintln(output, "Thank you for your soup rating!")
	}
}

func (c *rateCommand) RequiresAdmin() bool {
	return false
}

func (c *rateCommand) Usage() string {
	return "`<@soupbot> rate <score 1 to 10> <optional comment>` Rate the soup of the day, optionally include a comment."
}

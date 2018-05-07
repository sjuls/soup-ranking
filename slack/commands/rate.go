package commands

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/sjuls/soup-ranking/soup"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	rateCommand struct {
		repo        dbctx.ScoreRepository
		soupManager *soup.Manager
	}
)

// NewRateCommand create a new rate command
func NewRateCommand(repo dbctx.ScoreRepository, soupManager *soup.Manager) Command {
	return &rateCommand{
		repo,
		soupManager,
	}
}

func (c *rateCommand) Execute(args string, output io.Writer) {
	rateMatcher, err := regexp.Compile("^\\s*(10|[1-9])\\s*(.*)$")
	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	rateMatches := rateMatcher.FindStringSubmatch(args)

	if len(rateMatches) < 2 {
		fmt.Fprintln(output, "Quit messing with me and GIVE ME THAT SOUP RATING!")
		return
	}

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
		score.Comment = &rateMatches[2]
	}

	soupName, err := c.soupManager.GetSoupName()
	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	err = c.repo.SaveScore(&score)

	if err != nil {
		fmt.Fprintln(output, err.Error())
		return
	}

	fmt.Fprintf(output, "You've rated the soup `%s` a %d!\n\n", *soupName, scoreValue)

	if len(args) > 2 {
		fmt.Fprintf(output, "You left the following comment:\n```%s```\n\n", rateMatches[2])
	}

	fmt.Fprintln(output, "Thank you for your soup rating!")
}

func (c *rateCommand) RequiresAdmin() bool {
	return false
}

func (c *rateCommand) Usage() string {
	return "`<@soupbot> rate <score 1 to 10> <optional comment>` Rate the soup of the day, optionally include a comment."
}

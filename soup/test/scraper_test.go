package soup_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sjuls/soup-ranking/soup"
)

type (
	menuResponder struct{}
)

func TestSoupScraper(t *testing.T) {
	responder := menuResponder{}

	soupScraper := soup.NewScraper(&responder)
	soupOfTheDay, err := soupScraper.GetSoupOfTheDayName()

	if err != nil {
		t.Error(err)
		return
	}

	if *soupOfTheDay != "Ærtesuppe med ristede kerner" {
		t.Error("Wrong soup of the day")
		return
	}
}

func (m *menuResponder) Do(req *http.Request) (*http.Response, error) {
	menuBytes, err := ioutil.ReadFile("menu.html")
	if err != nil {
		return nil, err
	}
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(menuBytes)),
	}, nil
}

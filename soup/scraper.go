package soup

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/antchfx/htmlquery"

	"github.com/sjuls/soup-ranking/utils"
)

const (
	menuURL           = "https://iss.inmsystems.com/TakeAway/DanskeBankEjby963/Main/Products/1"
	soupOfTheDayTitle = "Dagens suppe"
)

type (
	// Scraper scrapes the internet for soup data
	Scraper interface {
		GetSoupOfTheDayName() (*string, error)
	}

	soupScraper struct {
		httpClient utils.HTTPClient
	}
)

// NewScraper creates a new scraper which can be used to access soup data
func NewScraper(httpClient utils.HTTPClient) Scraper {
	return &soupScraper{
		httpClient,
	}
}

func (s *soupScraper) GetSoupOfTheDayName() (*string, error) {
	menuHTML, err := s.getMenuHTML()
	if err != nil {
		return nil, err
	}

	soupOfTheDay, err := s.extractSoupOfTheDay(menuHTML)
	if err != nil {
		return nil, err
	}

	return soupOfTheDay, nil
}

func (s *soupScraper) getMenuHTML() (io.Reader, error) {
	log.Println("Fetching menu")
	request, err := http.NewRequest("GET", menuURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (s *soupScraper) extractSoupOfTheDay(menuHTML io.Reader) (*string, error) {
	document, err := htmlquery.Parse(menuHTML)
	if err != nil {
		return nil, err
	}

	for _, node := range htmlquery.Find(document, "(//div[@class=\"product\"])/div[@class=\"description\"]") {
		titleNodes := htmlquery.Find(node, "//h3/text()")
		descriptionNodes := htmlquery.Find(node, "//p/text()")

		if len(titleNodes) != 1 && len(descriptionNodes) != 1 {
			return nil, errors.New("Unexpected result from parsing product descriptions")
		}

		title := titleNodes[0].Data
		description := descriptionNodes[0].Data

		if title == soupOfTheDayTitle {
			return &description, nil
		}
	}

	return nil, errors.New("Could not find the soup of the day")
}

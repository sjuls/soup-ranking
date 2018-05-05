package soup

import (
	"github.com/sjuls/soup-ranking/dbctx"
)

var (
	unknownSoupName = "Unknown soup"
)

type (
	// Manager incapsulates soup business logic and exposes methods to interact with soup data
	Manager struct {
		repo    dbctx.SoupRepository
		scraper Scraper
	}
)

// NewManager creates a new manager
func NewManager(repo dbctx.SoupRepository, scraper Scraper) *Manager {
	return &Manager{
		repo,
		scraper,
	}
}

// GetSoupName extracts the name of the soup of the day
func (m *Manager) GetSoupName() (*string, error) {
	soupOfTheDay, err := m.repo.GetSoupOfTheDay()
	if err != nil {
		return nil, err
	}

	if soupOfTheDay != nil {
		return soupOfTheDay.Soup.Name, nil
	}

	return m.SyncSoupName()
}

// SyncSoupName updates the soup name with the soup name in the menu
func (m *Manager) SyncSoupName() (*string, error) {
	soupName, err := m.scraper.GetSoupOfTheDayName()
	if err != nil {
		return nil, err
	}

	if soupName == nil {
		soupName = &unknownSoupName
	}

	soupOfTheDay, err := m.repo.SetSoupOfTheDay(*soupName)
	if err != nil {
		return nil, err
	}

	return soupOfTheDay.Soup.Name, nil
}

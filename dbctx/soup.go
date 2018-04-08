package dbctx

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	soupRepository struct {
		connFactory ConnectionFactory
	}
	// SoupRepository to provide access to methods to handle soup data
	SoupRepository interface {
		// SetSoup sets the soup of the day
		SetSoup(name string) error
		GetSoupOfTheDay() (*SoupOfTheDay, error)
	}
)

// NewSoupRepository creates a new repository
func NewSoupRepository(connFactory ConnectionFactory) SoupRepository {
	return &soupRepository{
		connFactory,
	}
}

func (m *soupRepository) SetSoup(name string) error {
	db, err := m.connFactory.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()

	soupOfTheDay := getCreateSoupOfTheDay(tx)
	soupOfTheDay.Soup = *getSoup(&name, tx)

	tx.Save(soupOfTheDay)

	tx.Commit()
	return tx.Error
}

func (m *soupRepository) GetSoupOfTheDay() (*SoupOfTheDay, error) {
	db, err := m.connFactory.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx := db.Begin()
	soupOfTheDay := getCreateSoupOfTheDay(tx)
	tx.Commit()
	return soupOfTheDay, tx.Error
}

// getCreateSoupOfTheDay returns the soup of the day, pass DB or transaction to be used to extract the data.
func getCreateSoupOfTheDay(tx *gorm.DB) *SoupOfTheDay {
	soupOfTheDay := SoupOfTheDay{}
	today := time.Now().Truncate(24 * time.Hour)
	tx.FirstOrCreate(&soupOfTheDay, SoupOfTheDay{Date: &today})

	if &soupOfTheDay.Soup == nil {
		soupOfTheDay.Soup = *getSoup(nil, tx)
	}

	return &soupOfTheDay
}

// getSoup get soup with the given name - gets the unknown soup if name is nil
func getSoup(name *string, tx *gorm.DB) *Soup {
	soup := Soup{}

	if name == nil {
		newName := "Unknown soup"
		name = &newName
	}

	tx.FirstOrCreate(&soup, Soup{Name: *name})
	return &soup
}

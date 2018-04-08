package soup

import (
	"time"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	soupRepository struct {
		connFactory dbctx.ConnectionFactory
	}
	// Repository to provide access to methods to handle soup data
	Repository interface {
		// SetSoup sets the soup of the day
		SetSoup(name string) error
	}
)

// NewRepository creates a new repository
func NewRepository(connFactory dbctx.ConnectionFactory) Repository {
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

	var soup = dbctx.Soup{}
	tx.FirstOrCreate(&soup, dbctx.Soup{Name: name})

	soupOfTheDay := &dbctx.SoupOfTheDay{}
	tx.Where("date=DATE(?)", time.Now()).FirstOrCreate(&soupOfTheDay, dbctx.SoupOfTheDay{})
	soupOfTheDay.Soup = &soup

	tx.Save(soup)
	tx.Save(soupOfTheDay)

	tx.Commit()
	return nil
}

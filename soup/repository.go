package soup

import (
	"time"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	// Repository to provide access to methods to handle soup data
	Repository struct {
	}
)

// SetSoup sets the soup of the day
func (m *Repository) SetSoup(name string) error {
	db, err := dbctx.Open()
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

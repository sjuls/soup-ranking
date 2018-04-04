package score

import (
	"time"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	// Manager provides access to methods to handle scores
	Repository struct {
	}
)

// GetScores retrieves the total combined scores.
// TODO: Remove and replace with something more usable.
func (m *Repository) GetScores() (*[]dbctx.Score, error) {
	db, err := dbctx.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var scores []dbctx.Score

	db.Find(&scores)

	return &scores, nil
}

// SaveScore saves a score to the soup of the day
func (m *Repository) SaveScore(score *dbctx.Score) error {
	db, err := dbctx.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()

	soupOfTheDay := &dbctx.SoupOfTheDay{}
	tx.Where("date=DATE(?)", time.Now()).FirstOrCreate(&soupOfTheDay, dbctx.SoupOfTheDay{})
	score.SoupOfTheDay = soupOfTheDay

	tx.Save(soupOfTheDay)
	tx.Create(score)

	tx.Commit()

	return nil
}

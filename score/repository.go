package score

import (
	"time"

	"github.com/sjuls/soup-ranking/dbctx"
)

type (
	scoreRepository struct {
		connFactory dbctx.ConnectionFactory
	}

	// Repository gives access to methods used to persist and query scores
	Repository interface {
		GetScores() (*[]dbctx.Score, error)
		SaveScore(score *dbctx.Score) error
	}
)

// NewRepository creates a new repository
func NewRepository(connFactory dbctx.ConnectionFactory) Repository {
	var repo Repository = &scoreRepository{
		connFactory,
	}
	return repo
}

// GetScores retrieves the total combined scores.
// TODO: Remove and replace with something more usable.
func (r *scoreRepository) GetScores() (*[]dbctx.Score, error) {
	db, err := r.connFactory.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var scores []dbctx.Score

	db.Find(&scores)

	return &scores, nil
}

// SaveScore saves a score to the soup of the day
func (r *scoreRepository) SaveScore(score *dbctx.Score) error {
	db, err := r.connFactory.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()

	soupOfTheDay := &dbctx.SoupOfTheDay{}
	tx.Where("day=DATE(?)", time.Now()).FirstOrCreate(&soupOfTheDay, dbctx.SoupOfTheDay{})
	score.SoupOfTheDay = soupOfTheDay

	tx.Save(soupOfTheDay)
	tx.Create(score)

	tx.Commit()

	return nil
}

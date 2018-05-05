package dbctx

type (
	scoreRepository struct {
		connFactory ConnectionFactory
	}

	// ScoreRepository gives access to methods used to persist and query scores
	ScoreRepository interface {
		GetScores() (*[]Score, error)
		SaveScore(score *Score) error
	}
)

// NewScoreRepository creates a new repository
func NewScoreRepository(connFactory ConnectionFactory) ScoreRepository {
	return &scoreRepository{
		connFactory,
	}
}

// GetScores retrieves the total combined scores.
// TODO: Remove and replace with something more usable.
func (r *scoreRepository) GetScores() (*[]Score, error) {
	db, err := r.connFactory.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var scores []Score

	db.Find(&scores)

	return &scores, db.Error
}

// SaveScore saves a score to the soup of the day
func (r *scoreRepository) SaveScore(score *Score) error {
	db, err := r.connFactory.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()

	soupOfTheDay := getCreateSoupOfTheDay(tx)
	score.SoupOfTheDay = soupOfTheDay

	tx.Create(score)
	tx.Commit()

	return tx.Error
}

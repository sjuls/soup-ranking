package dbctx

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	// Soup holds soup metadata
	Soup struct {
		gorm.Model
		Name          string `gorm:"unique_index"`
		SoupOfTheDays []SoupOfTheDay
	}

	// SoupOfTheDay holds information regarding the soup of the day
	SoupOfTheDay struct {
		gorm.Model
		Date   *time.Time `gorm:"unique_index"`
		SoupID uint
		Soup   Soup
		Scores []Score
	}

	// Score holds information submitted by users regarding the soup of the day
	Score struct {
		gorm.Model
		Score          int
		Comment        string
		SoupOfTheDayID uint
		SoupOfTheDay   SoupOfTheDay `gorm:"not_null"`
	}
)

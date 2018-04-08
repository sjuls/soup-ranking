package dbctx

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	// Soup holds soup metadata
	Soup struct {
		gorm.Model
		Name string `gorm:"unique_index"`
	}

	// SoupOfTheDay holds information regarding the soup of the day
	SoupOfTheDay struct {
		gorm.Model
		day  time.Time `sql:"type:date;DEFAULT:current_date"`
		Soup *Soup
	}

	// Score holds information submitted by users regarding the soup of the day
	Score struct {
		gorm.Model
		Score        int
		Comment      string
		SoupOfTheDay *SoupOfTheDay
	}
)

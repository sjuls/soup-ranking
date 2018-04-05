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

	// SoupOfTheDay holds information rega
	SoupOfTheDay struct {
		gorm.Model
		date time.Time `sql:"type:date;unique_index;DEFAULT:current_date"`
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

package dbctx

import (
	"fmt"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import postgres dialect for GORM.
)

var (
	databaseURL string
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

// Open returns an open database connection.
func Open() (db *gorm.DB, err error) {
	return gorm.Open("postgres", databaseURL)
}

// Init - call to migrate the database and enable the use of the Open function.
func Init(database *string) error {
	var err error
	databaseURL, err = normalizeDatabaseURL(database)
	if err != nil {
		return err
	}

	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&Soup{}, &SoupOfTheDay{}, &Score{})

	return nil
}

func normalizeDatabaseURL(databaseURL *string) (string, error) {
	regex, err := regexp.Compile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	if err != nil {
		return "", err
	}

	match := regex.FindStringSubmatch(*databaseURL)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", match[3], match[4], match[1], match[2], match[5]), nil
}

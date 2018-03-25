package dbctx

import (
	"fmt"
	"regexp"

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
		Name string
	}

	// Score holds information submitted by users regarding the soup of the day
	Score struct {
		gorm.Model
		Score   int
		Comment string
		Soup    *Soup
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

	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&Soup{}, &Score{})

	return nil
}

func normalizeDatabaseURL(databaseURL *string) (string, error) {
	r, err := regexp.Compile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	if err != nil {
		return "", err
	}

	match := r.FindStringSubmatch(*databaseURL)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", match[3], match[4], match[1], match[2], match[5]), nil
}

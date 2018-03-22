package dbctx

import (
	"fmt"
	"regexp"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	databaseURL string
)

type (
	Soup struct {
		gorm.Model
		Name    	string
	}

	Score struct {
		gorm.Model
		Score		 	int
		Soup   		*Soup
	}
)

func Open() (db *gorm.DB, err error) {
	return gorm.Open("postgres", databaseURL)
}

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
	
	db.AutoMigrate(&Soup{},&Score{})

	return nil
}

func normalizeDatabaseURL(databaseURL *string) (string, error) {
	r, err := regexp.Compile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	if err != nil {
		return "", err
	}

	match := r.FindStringSubmatch(*databaseURL)

	return	fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", match[3], match[4], match[1], match[2], match[5]), nil
}
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
	connectionFactory struct {
		databaseURL string
	}

	// ConnectionFactory encapsulates gorm database access
	ConnectionFactory interface {
		Open() (db *gorm.DB, err error)
	}
)

// NewConnectionFactory is used to create a factory to create DB connections
func NewConnectionFactory(connectionString *string) (ConnectionFactory, error) {
	var err error
	databaseURL, err = normalizeDatabaseURL(connectionString)
	if err != nil {
		return nil, err
	}

	connection, err := gorm.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	defer connection.Close()

	connection.AutoMigrate(&Soup{}, &SoupOfTheDay{}, &Score{})

	var connFactory ConnectionFactory = &connectionFactory{
		databaseURL: databaseURL,
	}

	return connFactory, nil
}

// Open returns an open database connection.
func (cf *connectionFactory) Open() (db *gorm.DB, err error) {
	return gorm.Open("postgres", cf.databaseURL)
}

func normalizeDatabaseURL(databaseURL *string) (string, error) {
	regex, err := regexp.Compile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	if err != nil {
		return "", err
	}

	match := regex.FindStringSubmatch(*databaseURL)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", match[3], match[4], match[1], match[2], match[5]), nil
}

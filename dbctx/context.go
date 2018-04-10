package dbctx

import (
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import postgres dialect for GORM.
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // Import sqlit dialect for GORM.
)

type (
	connectionFactory struct {
		connectionBuilder func() (*gorm.DB, error)
	}

	// ConnectionFactory encapsulates gorm database access
	ConnectionFactory interface {
		Open() (db *gorm.DB, err error)
	}
)

// NewConnectionFactory is used to create a factory to create DB connections
func NewConnectionFactory(connectionString string) (ConnectionFactory, error) {
	var connectionFactory ConnectionFactory
	var err error
	if len(connectionString) == 0 {
		connectionFactory = newSQLiteConnectionFactory()
	} else {
		connectionFactory, err = newPostgresConnectionFactory(&connectionString)
		if err != nil {
			return nil, err
		}
	}

	connection, err := connectionFactory.Open()
	if err != nil {
		return nil, err
	}

	connection.AutoMigrate(&Soup{}, &SoupOfTheDay{}, &Score{})

	return connectionFactory, nil
}

// Open returns an open database connection.
func (cf *connectionFactory) Open() (db *gorm.DB, err error) {
	return cf.connectionBuilder()
}

func newPostgresConnectionFactory(connectionString *string) (ConnectionFactory, error) {
	databaseURL, err := normalizeDatabaseURL(connectionString)
	if err != nil {
		return nil, err
	}

	connectionBuilder := func() (*gorm.DB, error) {
		return gorm.Open("postgres", databaseURL)
	}

	return newConnectionFactory(connectionBuilder), nil
}

func newSQLiteConnectionFactory() ConnectionFactory {
	connectionBuilder := func() (*gorm.DB, error) {
		return gorm.Open("sqlite3", "/gorm.db")
	}

	return newConnectionFactory(connectionBuilder)
}

func newConnectionFactory(connectionBuilder func() (*gorm.DB, error)) ConnectionFactory {
	return &connectionFactory{
		connectionBuilder: connectionBuilder,
	}
}

func normalizeDatabaseURL(databaseURL *string) (string, error) {
	regex, err := regexp.Compile("^postgres://(.+):(.+)@(.+):(.+)/(.+)$")
	if err != nil {
		return "", err
	}

	match := regex.FindStringSubmatch(*databaseURL)

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", match[3], match[4], match[1], match[2], match[5]), nil
}

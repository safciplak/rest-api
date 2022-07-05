package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// NewDatabase - returns a pointer to a database object
func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting up new database connection")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUsername, dbName, dbPassword, sslMode)

	db, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, err
}

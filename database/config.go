package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// NewDB creates a new database connection and returns the GORM instance.
func NewDB() (*gorm.DB, error) {
	err := godotenv.Load("../config/db.env")
	if err != nil {
		// Handle the error if the .env file is not found or there's an error reading it.
		panic("Error loading .env file: " + err.Error())
	}

	// Get the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}

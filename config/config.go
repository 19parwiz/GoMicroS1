package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDatabase initializes and returns a PostgreSQL DB connection
func InitDatabase() (*gorm.DB, error) {
	// You can set these as environment variables or hardcode during testing
	host := "localhost"
	user := "postgres"
	password := "54321"
	dbname := "ecomventory"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	return db, nil
}

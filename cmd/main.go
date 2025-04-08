package main

import (
	"ecomventory/config"
	"ecomventory/router"
	"log"
)

func main() {
	// Initialize the database connection
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return
	}

	// Setup the router and pass the db connection
	r := router.SetupRouter(db) // Passing db to the router

	// Start the server
	r.Run(":8080")
}

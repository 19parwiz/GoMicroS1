package main

import (
	"ecomventory/config"
	"ecomventory/router"
	"log"
)

func main() {
	// Connect to the database using the ConnectDB method from config package
	db := config.ConnectDB()

	// Use db to initialize the router
	r := router.SetupRouter(db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

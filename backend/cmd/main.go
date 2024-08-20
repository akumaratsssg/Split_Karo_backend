package main

import (
	"Expense_Management/backend/config"
	"Expense_Management/backend/database"
	"Expense_Management/backend/routes"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Connect to the database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Setup the router and routes
	r := routes.SetupRoutes(db)

	// Set trusted proxies (example: only trust localhost)
	// Adjust the list of trusted proxies according to your deployment environment
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Could not set trusted proxies: %v", err)
	}

	// Start the server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"

	"sociomile-apps/config"
	"sociomile-apps/internal/database"
	"sociomile-apps/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting Bill Tracker API in %s mode", cfg.DBHost)

	db := database.Connect(cfg)

	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := routes.SetupRouter(db)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

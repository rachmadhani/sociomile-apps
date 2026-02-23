package main

import (
	"log"

	"sociomile-apps/config"
	"sociomile-apps/internal/database"
	"sociomile-apps/internal/event"
	"sociomile-apps/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting Bill Tracker API in %s mode", cfg.DBHost)

	db := database.Connect(cfg)
	dispatcher := event.NewDispatcher(100)
	worker := event.NewWorker(db)
	worker.Start(dispatcher.Channel())

	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := routes.SetupRouter(db, dispatcher)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

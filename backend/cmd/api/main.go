package main

import (
	"log"
	"os"

	"sociomile-apps/config"
	"sociomile-apps/internal/cache"
	"sociomile-apps/internal/database"
	"sociomile-apps/internal/event"
	"sociomile-apps/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Starting Sociomile API in %s mode", cfg.DBHost)

	db := database.Connect(cfg)
	dispatcher := event.NewDispatcher(100)
	worker := event.NewWorker(db)
	worker.Start(dispatcher.Channel())

	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	redisClient := cache.NewRedisClient()

	conversationCache := cache.NewConversationCache(redisClient)
	ticketCache := cache.NewTicketCache(redisClient)

	router := routes.SetupRouter(db, dispatcher, conversationCache, ticketCache)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

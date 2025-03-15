package main

import (
	"log"

	"github.com/messaging-go-service/config"
	"github.com/messaging-go-service/internal/model"
)

func main() {
	config.OpenConnection()

	if config.Database == nil {
		log.Fatal("Database connection is nil")
	}

	if err := config.Database.AutoMigrate(
		&model.User{},
		&model.Conversation{},
		&model.Participant{},
		&model.Notification{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	log.Println("Database migration completed successfully.")
}

package database

import (
	"fmt"
	"log"

	"sociomile-apps/config"
	model "sociomile-apps/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return DB
}

func AutoMigrate() error {
	log.Println("Running database migrations...")
	err := DB.AutoMigrate(
		&model.User{},
		&model.Tenant{},
		&model.Conversation{},
		&model.Message{},
		&model.Ticket{},
	)
	if err != nil {
		return err
	}
	return nil
}

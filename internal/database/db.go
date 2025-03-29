package database

import (
	"fmt"

	"github.com/FJericho/test-mnc/internal/config"
	"github.com/FJericho/test-mnc/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func ConnectDB() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	db.Debug().AutoMigrate(model.User{}, model.History{}, model.Payment{})

	return db, nil
}
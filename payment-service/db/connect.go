package db

import (
	"fmt"
	"log"

	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/config"
	"github.com/JoaoDallagnol/go-restaurant-orders/payment-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg := config.AppConfig.Database
	schema := cfg.Schema

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode, schema,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get raw DB object: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	if err := DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error; err != nil {
		log.Fatalf("❌ Failed to create schema '%s': %v", schema, err)
	}

	if err := DB.AutoMigrate(&model.Payment{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("✅ Database connection successful")
}

func GetDB() *gorm.DB {
	return DB
}

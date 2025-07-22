package db

import (
	"fmt"
	"log"

	"github.com/JoaoDallagnol/go-restaurant-orders/auth-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Define your custom schema
	const schema = "go_restaurant_orders_auth_service"

	// DSN with search_path pointing to your schema
	dsn := fmt.Sprintf(
		"host=localhost user=admin password=admin123 dbname=postgres port=5432 sslmode=disable search_path=%s",
		schema,
	)

	// Open the GORM connection
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Get underlying *sql.DB for lower-level control (e.g., Ping)
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get raw DB object: %v", err)
	}

	// Ping the database to verify connectivity
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	// Create the schema if it doesn't exist
	if err := DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error; err != nil {
		log.Fatalf("❌ Failed to create schema '%s': %v", schema, err)
	}

	// Auto migrate schema
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("✅ Database connection successful")
}

func GetDB() *gorm.DB {
	return DB
}

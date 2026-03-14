package database

import (
	"fmt"
	"log"
	"ocra/pkg/entities"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Setup() *gorm.DB {

	supabaseConnectionString := os.Getenv("DATABASE_URL")
	if supabaseConnectionString == "" {
		log.Fatal("DATABASE_URL must be set in the .env file")

	}

	dbSchema := os.Getenv("DB_SCHEMA")
	if dbSchema == "" {
		log.Fatal("DB_SCHEMA must be set in the .env file")
	}

	db, err := gorm.Open(postgres.Open(supabaseConnectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbSchema + ".",
		},
	})
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}

	fmt.Println("Connected to Supabase successfully")

	synchronizeDB := os.Getenv("SYNCHRONIZE_DB")
	if synchronizeDB == "true" {
		fmt.Println("Synchronizing database")
		if err := db.AutoMigrate(
			&entities.User{},
			&entities.Zone{},
			&entities.Client{},
			&entities.House{},
			&entities.SearchHouse{},
		); err != nil {
			log.Fatalf("failed to synchronize database: %v", err)
		}
		fmt.Println("Synchronization complete")
	}
	return db
}

package commands

import (
	"log"

	"svc-dynamic-form-go/configs"
	_"svc-dynamic-form-go/src/models"
)

func Migrate() {
	log.Println("📦 Running migration...")

	// Load .env dan DB
	configs.LoadEnv()
	configs.InitDB()

	// Migrasikan semua model
	err := configs.DB.AutoMigrate()

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully.")
}

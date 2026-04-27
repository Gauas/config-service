package model

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&Config{},
	); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	log.Println("database migrated")
	return nil
}

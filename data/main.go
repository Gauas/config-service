package data

import (
	"log"

	"github.com/gauas/config-service/infra"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

func New(infraInstance *infra.Infra) *Data {
	dataInstance := &Data{db: infraInstance.DB}
	dataInstance.migrate()
	return dataInstance
}

func (d *Data) migrate() {
	if err := d.db.AutoMigrate(&Config{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("database migrated")
}

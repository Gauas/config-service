package data

import (
	"log"

	"github.com/gauas/config-service/infra"
	"gorm.io/gorm"
)

type Repositories struct {
	Config Repository[Config]
}

type Data struct {
	db         *gorm.DB
	Repository Repositories
}

func New(infraInstance *infra.Infra) *Data {
	d := &Data{db: infraInstance.DB}
	d.Repository = Repositories{
		Config: d.NewConfigRepo(),
	}
	d.migrate()
	return d
}

func (d *Data) migrate() {
	if err := d.db.AutoMigrate(&Config{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	log.Println("database migrated")
}

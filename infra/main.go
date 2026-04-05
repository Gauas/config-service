package infra

import (
	"github.com/gauas/config-service/config"
	"gorm.io/gorm"
)

type Infra struct {
	DB *gorm.DB
}

func New(appConfig config.AppConfig) *Infra {
	db := connectDatabase(appConfig.DBUrl)
	return &Infra{DB: db}
}

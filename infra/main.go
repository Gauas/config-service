package infra

import (
	"github.com/gauas/config-service/config"
	"gorm.io/gorm"
)

type Infra struct {
	DB *gorm.DB
}

func New(appConfig config.Config) *Infra {
	db := DatabaseConnection(appConfig.DBUrl)
	return &Infra{DB: db}
}

package repository

import "gorm.io/gorm"
import "github.com/gauas/config-service/model"

type Registry struct {
	Config Repository[model.Config]
}

func New(db *gorm.DB) *Registry {
	return &Registry{
		Config: Repository[model.Config]{db: db},
	}
}

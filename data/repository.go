package data

import (
	"errors"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	FindOne(query string, args ...any) (*T, error)
	Create(entity *T) error
	Save(entity *T) error
	Delete(query string, args ...any) error
}

type BaseRepo[T any] struct {
	db *gorm.DB
}

func (r *BaseRepo[T]) FindOne(query string, args ...any) (*T, error) {
	var entity T
	err := r.db.Where(query, args...).First(&entity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepo[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepo[T]) Save(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepo[T]) Delete(query string, args ...any) error {
	var entity T
	return r.db.Where(query, args...).Delete(&entity).Error
}

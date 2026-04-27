package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JSONMap map[string]interface{}

type Config struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Service     string         `gorm:"not null;uniqueIndex:uq_service_env"            json:"service"`
	Environment string         `gorm:"not null;uniqueIndex:uq_service_env"            json:"environment"`
	Config      JSONMap        `gorm:"type:jsonb;not null"                            json:"config"`
	CreatedAt   time.Time      `                                                      json:"created_at"`
	UpdatedAt   time.Time      `                                                      json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"uniqueIndex:uq_service_env;index"               json:"-"`
}

func (c *Config) Merge(payload JSONMap) error {
	if err := validatePayload(payload); err != nil {
		return err
	}
	if c.Config == nil {
		c.Config = JSONMap{}
	}
	for key, value := range payload {
		c.Config[key] = value
	}
	return nil
}

func validatePayload(payload JSONMap) error {
	if payload == nil {
		return fmt.Errorf("config is required")
	}
	for key, value := range payload {
		switch value.(type) {
		case map[string]interface{}:
			return fmt.Errorf("nested objects not allowed (key: %q)", key)
		case []interface{}:
			return fmt.Errorf("arrays not allowed (key: %q)", key)
		}
	}
	return nil
}

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

func (j *JSONMap) Scan(src interface{}) error {
	var bytes []byte
	switch value := src.(type) {
	case []byte:
		bytes = value
	case string:
		bytes = []byte(value)
	default:
		return errors.New("unsupported type for JSONMap")
	}
	return json.Unmarshal(bytes, j)
}

package data

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JSONMap map[string]any

type Config struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Service     string         `gorm:"not null;uniqueIndex:uq_service_env"            json:"service"`
	Environment string         `gorm:"not null;uniqueIndex:uq_service_env"            json:"environment"`
	Config      JSONMap        `gorm:"type:jsonb;not null"                            json:"config"`
	CreatedAt   time.Time      `                                                      json:"created_at"`
	UpdatedAt   time.Time      `                                                      json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"uniqueIndex:uq_service_env;index"               json:"-"`
}

type ConfigRepo struct {
	BaseRepo[Config]
}

func (d *Data) NewConfigRepo() Repository[Config] {
	return &ConfigRepo{BaseRepo: BaseRepo[Config]{db: d.db}}
}

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	bytes, err := json.Marshal(j)
	return string(bytes), err
}

func (j *JSONMap) Scan(src any) error {
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

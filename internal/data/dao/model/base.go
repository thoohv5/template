package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        int32          `gorm:"column:id;primaryKey" redis:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" redis:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" redis:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" redis:"-" json:"-"`
}

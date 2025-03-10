package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	IsDefault bool      `gorm:"not null"`
	Icon      string    `gorm:"null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *Category) TableName() string {
	return "categories"
}

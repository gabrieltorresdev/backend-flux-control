package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `gorm:"primaryKey"`
	KeycloakID string         `gorm:"not null;unique"`
	Name       string         `gorm:"not null"`
	Email      string         `gorm:"not null;unique"`
	Username   string         `gorm:"not null;unique"`
	Status     string         `gorm:"not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

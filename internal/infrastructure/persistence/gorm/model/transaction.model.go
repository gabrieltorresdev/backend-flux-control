package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	CategoryID  uuid.UUID `gorm:"not null"`
	UserID      uuid.UUID `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Datetime    time.Time `gorm:"not null"`
	Description string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}

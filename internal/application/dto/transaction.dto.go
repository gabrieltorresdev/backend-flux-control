package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTransactionDTO struct {
	UserID      uuid.UUID
	CategoryID  uuid.UUID
	Amount      float64
	Datetime    time.Time
	Description string
}

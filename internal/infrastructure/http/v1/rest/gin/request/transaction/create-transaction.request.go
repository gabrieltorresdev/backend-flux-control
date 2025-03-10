package transaction

import (
	"time"

	"github.com/gabrieltorresdev/backend-flux-control/internal/application/dto"
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	CategoryID  uuid.UUID
	Amount      float64
	Datetime    time.Time
	Description string
}

func (r *CreateTransactionRequest) ToCreateTransactionDTO(userId uuid.UUID) *dto.CreateTransactionDTO {
	return &dto.CreateTransactionDTO{
		UserID:      userId,
		CategoryID:  r.CategoryID,
		Amount:      r.Amount,
		Datetime:    r.Datetime,
		Description: r.Description,
	}
}

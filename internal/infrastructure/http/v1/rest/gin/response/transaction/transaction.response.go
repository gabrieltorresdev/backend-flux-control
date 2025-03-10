package transaction

import (
	"time"

	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity"
	"github.com/gabrieltorresdev/backend-flux-control/pkg/hateoas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"categoryId"`
	UserID      uuid.UUID `json:"userId"`
	Amount      float64   `json:"amount"`
	Datetime    time.Time `json:"datetime"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromEntity(t entity.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:          t.ID(),
		CategoryID:  t.CategoryID(),
		UserID:      t.UserID(),
		Amount:      t.Amount(),
		Datetime:    t.Datetime(),
		Description: t.Description(),
		CreatedAt:   t.CreatedAt(),
		UpdatedAt:   t.UpdatedAt(),
	}
}

func FromEntities(transactions []entity.Transaction) []TransactionResponse {
	result := make([]TransactionResponse, len(transactions))
	for i, transaction := range transactions {
		result[i] = FromEntity(transaction)
	}
	return result
}

func BuildTransactionResponse(ctx *gin.Context, transaction entity.Transaction, statusCode int) *hateoas.Response {
	transactionResponse := FromEntity(transaction)

	return hateoas.Single("transaction", transactionResponse, ctx, statusCode)
}

func BuildTransactionsResponse(ctx *gin.Context, transactions []entity.Transaction, page, pageSize int, statusCode int) *hateoas.Response {
	transactionsResponse := FromEntities(transactions)

	return hateoas.Collection("transaction", transactionsResponse, ctx, page, pageSize, len(transactions), statusCode)
}

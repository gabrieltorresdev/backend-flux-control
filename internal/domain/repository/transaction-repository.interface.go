package repository

import (
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/pagination"
)

type TransactionRepositoryInterface interface {
	FindAllPaginated(paginate *pagination.Pagination) ([]entity.Transaction, error)
	Create(transaction *entity.Transaction) (*entity.Transaction, error)
}

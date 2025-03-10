package interfaces

import (
	"github.com/gabrieltorresdev/backend-flux-control/internal/application/dto"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/pagination"
)

type TransactionServiceInterface interface {
	FindAllPaginated(page, pageSize int) ([]entity.Transaction, *pagination.Pagination, error)
	Create(createTransactionDTO *dto.CreateTransactionDTO) (*entity.Transaction, error)
}

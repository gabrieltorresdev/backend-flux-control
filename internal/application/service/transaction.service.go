package service

import (
	"time"

	"github.com/gabrieltorresdev/backend-flux-control/internal/application/dto"
	"github.com/gabrieltorresdev/backend-flux-control/internal/application/interfaces"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/pagination"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/repository"
	"github.com/google/uuid"
)

type TransactionService struct {
	transactionRepository repository.TransactionRepositoryInterface
}

func NewTransactionService(transactionRepository repository.TransactionRepositoryInterface) interfaces.TransactionServiceInterface {
	return &TransactionService{transactionRepository: transactionRepository}
}

func (s *TransactionService) FindAllPaginated(page, pageSize int) ([]entity.Transaction, *pagination.Pagination, error) {
	paginate := pagination.NewPagination(page, pageSize)

	transactions, err := s.transactionRepository.FindAllPaginated(paginate)
	if err != nil {
		return nil, nil, err
	}

	return transactions, paginate, nil
}

func (s *TransactionService) Create(createTransactionDTO *dto.CreateTransactionDTO) (*entity.Transaction, error) {
	transaction, err := entity.NewTransaction(
		uuid.New(),
		createTransactionDTO.CategoryID,
		createTransactionDTO.UserID,
		createTransactionDTO.Amount,
		createTransactionDTO.Datetime,
		createTransactionDTO.Description,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	createdTransaction, err := s.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

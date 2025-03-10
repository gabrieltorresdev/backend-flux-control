package repository

import (
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/pagination"
	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/repository"
	"github.com/gabrieltorresdev/backend-flux-control/internal/infrastructure/persistence/gorm/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	gorm *gorm.DB
}

func NewTransactionRepository(gorm *gorm.DB) repository.TransactionRepositoryInterface {
	return &TransactionRepository{gorm: gorm}
}

func (r *TransactionRepository) FindAllPaginated(paginate *pagination.Pagination) ([]entity.Transaction, error) {
	var transactions []model.Transaction
	var totalItems int64

	if err := r.gorm.Model(&model.Transaction{}).Count(&totalItems).Error; err != nil {
		return nil, err
	}

	paginate.SetTotal(totalItems)

	if err := r.gorm.Offset(paginate.GetOffset()).Limit(paginate.GetLimit()).Find(&transactions).Error; err != nil {
		return nil, err
	}

	transactionsEntity := make([]entity.Transaction, len(transactions))
	for i, transaction := range transactions {
		var err error
		transactionEntity, err := entity.NewTransaction(
			transaction.ID,
			transaction.CategoryID,
			transaction.UserID,
			transaction.Amount,
			transaction.Datetime,
			transaction.Description,
			transaction.CreatedAt,
			transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactionsEntity[i] = *transactionEntity
	}

	return transactionsEntity, nil
}

func (r *TransactionRepository) Create(transaction *entity.Transaction) (*entity.Transaction, error) {
	transactionModel := model.Transaction{
		ID:          transaction.ID(),
		CategoryID:  transaction.CategoryID(),
		UserID:      transaction.UserID(),
		Amount:      transaction.Amount(),
		Datetime:    transaction.Datetime(),
		Description: transaction.Description(),
		CreatedAt:   transaction.CreatedAt(),
		UpdatedAt:   transaction.UpdatedAt(),
	}
	if err := r.gorm.Create(&transactionModel).Error; err != nil {
		return nil, err
	}

	transactionEntity, err := entity.NewTransaction(
		transactionModel.ID,
		transactionModel.CategoryID,
		transactionModel.UserID,
		transactionModel.Amount,
		transactionModel.Datetime,
		transactionModel.Description,
		transactionModel.CreatedAt,
		transactionModel.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return transactionEntity, nil
}

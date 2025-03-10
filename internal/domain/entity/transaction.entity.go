package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	id          uuid.UUID
	categoryID  uuid.UUID
	userID      uuid.UUID
	amount      float64
	datetime    time.Time
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *Transaction) ID() uuid.UUID         { return t.id }
func (t *Transaction) CategoryID() uuid.UUID { return t.categoryID }
func (t *Transaction) UserID() uuid.UUID     { return t.userID }
func (t *Transaction) Amount() float64       { return t.amount }
func (t *Transaction) Datetime() time.Time   { return t.datetime }
func (t *Transaction) Description() string   { return t.description }
func (t *Transaction) CreatedAt() time.Time  { return t.createdAt }
func (t *Transaction) UpdatedAt() time.Time  { return t.updatedAt }

func NewTransaction(id uuid.UUID, categoryID uuid.UUID, userID uuid.UUID, amount float64, datetime time.Time, description string, createdAt time.Time, updatedAt time.Time) (*Transaction, error) {
	transaction := &Transaction{
		id:          id,
		categoryID:  categoryID,
		userID:      userID,
		amount:      amount,
		datetime:    datetime,
		description: description,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}

	err := transaction.validate()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) validate() error {
	if t.categoryID == uuid.Nil {
		return errors.New("category id is required")
	}

	if t.userID == uuid.Nil {
		return errors.New("user id is required")
	}

	if t.amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if t.datetime.IsZero() {
		return errors.New("datetime is required")
	}

	return nil
}

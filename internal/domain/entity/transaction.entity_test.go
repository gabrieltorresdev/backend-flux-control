package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	t.Run("should create a new transaction successfully", func(t *testing.T) {
		categoryID := uuid.New()
		userID := uuid.New()
		amount := 100.0
		datetime := time.Now()
		description := "Grocery shopping"

		transaction, err := NewTransaction(categoryID, userID, uuid.Nil, amount, datetime, description, time.Time{}, time.Time{})

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.NotEqual(t, uuid.Nil, transaction.ID())
		assert.Equal(t, categoryID, transaction.CategoryID())
		assert.Equal(t, userID, transaction.UserID())
		assert.Equal(t, amount, transaction.Amount())
		assert.Equal(t, datetime, transaction.Datetime())
		assert.Equal(t, description, transaction.Description())
		assert.False(t, transaction.CreatedAt().IsZero())
		assert.False(t, transaction.UpdatedAt().IsZero())
	})

	t.Run("should return error when category id is not provided", func(t *testing.T) {
		userID := uuid.New()
		transaction, err := NewTransaction(uuid.Nil, userID, uuid.Nil, 100.0, time.Now(), "description", time.Time{}, time.Time{})

		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, "category id is required", err.Error())
	})

	t.Run("should return error when user id is not provided", func(t *testing.T) {
		categoryID := uuid.New()
		transaction, err := NewTransaction(categoryID, uuid.Nil, uuid.Nil, 100.0, time.Now(), "description", time.Time{}, time.Time{})

		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, "user id is required", err.Error())
	})

	t.Run("should return error when amount is zero or negative", func(t *testing.T) {
		categoryID := uuid.New()
		userID := uuid.New()

		transaction, err := NewTransaction(categoryID, userID, uuid.Nil, 0, time.Now(), "description", time.Time{}, time.Time{})
		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, "amount must be greater than 0", err.Error())

		transaction, err = NewTransaction(categoryID, userID, uuid.Nil, -10.0, time.Now(), "description", time.Time{}, time.Time{})
		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, "amount must be greater than 0", err.Error())
	})

	t.Run("should return error when datetime is zero", func(t *testing.T) {
		categoryID := uuid.New()
		userID := uuid.New()
		transaction, err := NewTransaction(categoryID, userID, uuid.Nil, 100.0, time.Time{}, "description", time.Time{}, time.Time{})

		assert.NotNil(t, err)
		assert.Nil(t, transaction)
		assert.Equal(t, "datetime is required", err.Error())
	})

	t.Run("should create transaction with empty description", func(t *testing.T) {
		categoryID := uuid.New()
		userID := uuid.New()
		transaction, err := NewTransaction(categoryID, userID, uuid.Nil, 100.0, time.Now(), "", time.Time{}, time.Time{})

		assert.Nil(t, err)
		assert.NotNil(t, transaction)
		assert.Equal(t, "", transaction.Description())
	})
}

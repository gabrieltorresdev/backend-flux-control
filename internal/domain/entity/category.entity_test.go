package entity

import (
	"testing"

	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity/enum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	t.Run("should create a new category successfully", func(t *testing.T) {
		userID := uuid.New()
		category, err := NewCategory(userID, "Food", enum.CategoryTypeExpense, false, "food-icon")

		assert.Nil(t, err)
		assert.NotNil(t, category)
		assert.NotEqual(t, uuid.Nil, category.ID())
		assert.Equal(t, userID, category.UserID())
		assert.Equal(t, "Food", category.Name())
		assert.Equal(t, enum.CategoryTypeExpense, category.Type())
		assert.Equal(t, false, category.Default())
		assert.Equal(t, "food-icon", category.Icon())
		assert.False(t, category.CreatedAt().IsZero())
		assert.False(t, category.UpdatedAt().IsZero())
	})

	t.Run("should return error when user id is not provided", func(t *testing.T) {
		category, err := NewCategory(uuid.Nil, "Food", enum.CategoryTypeExpense, false, "food-icon")

		assert.NotNil(t, err)
		assert.Nil(t, category)
		assert.Equal(t, "user id is required", err.Error())
	})

	t.Run("should return error when name is not provided", func(t *testing.T) {
		userID := uuid.New()
		category, err := NewCategory(userID, "", enum.CategoryTypeExpense, false, "food-icon")

		assert.NotNil(t, err)
		assert.Nil(t, category)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when type is invalid", func(t *testing.T) {
		userID := uuid.New()
		category, err := NewCategory(userID, "Food", "invalid", false, "food-icon")

		assert.NotNil(t, err)
		assert.Nil(t, category)
		assert.Equal(t, "invalid type", err.Error())
	})

	t.Run("should create category with default values", func(t *testing.T) {
		userID := uuid.New()
		category, err := NewCategory(userID, "Salary", enum.CategoryTypeIncome, true, "salary-icon")

		assert.Nil(t, err)
		assert.NotNil(t, category)
		assert.Equal(t, true, category.Default())
		assert.Equal(t, enum.CategoryTypeIncome, category.Type())
	})
}

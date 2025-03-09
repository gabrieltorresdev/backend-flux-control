package entity

import (
	"errors"
	"time"

	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity/enum"
	"github.com/google/uuid"
)

type Category struct {
	id              uuid.UUID
	userID          uuid.UUID
	name            string
	typeCategory    enum.CategoryType
	defaultCategory bool
	icon            string
	createdAt       time.Time
	updatedAt       time.Time
}

func (c *Category) ID() uuid.UUID           { return c.id }
func (c *Category) UserID() uuid.UUID       { return c.userID }
func (c *Category) Name() string            { return c.name }
func (c *Category) Type() enum.CategoryType { return c.typeCategory }
func (c *Category) Default() bool           { return c.defaultCategory }
func (c *Category) Icon() string            { return c.icon }
func (c *Category) CreatedAt() time.Time    { return c.createdAt }
func (c *Category) UpdatedAt() time.Time    { return c.updatedAt }

func NewCategory(userID uuid.UUID, name string, typeCategory enum.CategoryType, defaultCategory bool, icon string) (*Category, error) {
	category := &Category{
		id:              uuid.New(),
		userID:          userID,
		name:            name,
		typeCategory:    typeCategory,
		defaultCategory: defaultCategory,
		icon:            icon,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}

	err := category.validate()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) validate() error {
	if c.userID == uuid.Nil {
		return errors.New("user id is required")
	}

	if c.name == "" {
		return errors.New("name is required")
	}

	if c.typeCategory != enum.CategoryTypeIncome && c.typeCategory != enum.CategoryTypeExpense {
		return errors.New("invalid type")
	}

	return nil
}

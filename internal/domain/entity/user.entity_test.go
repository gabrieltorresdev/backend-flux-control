package entity

import (
	"testing"

	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity/enum"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("should create a new user successfully", func(t *testing.T) {
		user, err := NewUser("keycloak-123", "John Doe", "john@example.com", "johndoe", enum.UserStatusActive)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEqual(t, uuid.Nil, user.ID())
		assert.Equal(t, "keycloak-123", user.KeycloakID())
		assert.Equal(t, "John Doe", user.Name())
		assert.Equal(t, "john@example.com", user.Email())
		assert.Equal(t, "johndoe", user.Username())
		assert.Equal(t, enum.UserStatusActive, user.Status())
		assert.False(t, user.CreatedAt().IsZero())
		assert.False(t, user.UpdatedAt().IsZero())
	})

	t.Run("should return error when keycloak id is not provided", func(t *testing.T) {
		user, err := NewUser("", "John Doe", "john@example.com", "johndoe", enum.UserStatusActive)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "keycloak id is required", err.Error())
	})

	t.Run("should return error when keycloak id is not provided", func(t *testing.T) {
		user, err := NewUser("", "John Doe", "john@example.com", "johndoe", enum.UserStatusActive)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "keycloak id is required", err.Error())
	})

	t.Run("should return error when name is not provided", func(t *testing.T) {
		user, err := NewUser("keycloak-123", "", "john@example.com", "johndoe", enum.UserStatusActive)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return error when email is not provided", func(t *testing.T) {
		user, err := NewUser("keycloak-123", "John Doe", "", "johndoe", enum.UserStatusActive)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "email is required", err.Error())
	})

	t.Run("should return error when username is not provided", func(t *testing.T) {
		user, err := NewUser("keycloak-123", "John Doe", "john@example.com", "", enum.UserStatusActive)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "username is required", err.Error())
	})

	t.Run("should return error when status is invalid", func(t *testing.T) {
		user, err := NewUser("keycloak-123", "John Doe", "john@example.com", "johndoe", "invalid")

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid status", err.Error())
	})
}

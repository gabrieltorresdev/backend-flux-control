package entity

import (
	"errors"
	"time"

	"github.com/gabrieltorresdev/backend-flux-control/internal/domain/entity/enum"
	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	userID     uuid.UUID
	keycloakID string
	name       string
	email      string
	username   string
	status     enum.UserStatus
	createdAt  time.Time
	updatedAt  time.Time
}

func (u *User) ID() uuid.UUID           { return u.id }
func (u *User) UserID() uuid.UUID       { return u.userID }
func (u *User) KeycloakID() string      { return u.keycloakID }
func (u *User) Name() string            { return u.name }
func (u *User) Email() string           { return u.email }
func (u *User) Username() string        { return u.username }
func (u *User) Status() enum.UserStatus { return u.status }
func (u *User) CreatedAt() time.Time    { return u.createdAt }
func (u *User) UpdatedAt() time.Time    { return u.updatedAt }

func NewUser(userID uuid.UUID, keycloakID string, name string, email string, username string, status enum.UserStatus) (*User, error) {
	user := &User{
		id:         uuid.New(),
		userID:     userID,
		keycloakID: keycloakID,
		name:       name,
		email:      email,
		username:   username,
		status:     status,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}

	err := user.validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) validate() error {
	if u.userID == uuid.Nil {
		return errors.New("user id is required")
	}

	if u.keycloakID == "" {
		return errors.New("keycloak id is required")
	}

	if u.name == "" {
		return errors.New("name is required")
	}

	if u.email == "" {
		return errors.New("email is required")
	}

	if u.username == "" {
		return errors.New("username is required")
	}

	if u.status != enum.UserStatusPending && u.status != enum.UserStatusActive && u.status != enum.UserStatusInactive {
		return errors.New("invalid status")
	}

	return nil
}

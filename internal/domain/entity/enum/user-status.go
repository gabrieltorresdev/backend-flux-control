package enum

type UserStatus string

const (
	UserStatusPending  UserStatus = "pending"
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

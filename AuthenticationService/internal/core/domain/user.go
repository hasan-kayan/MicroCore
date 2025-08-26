package domain

import "time"

type UserID string

type UserStatus string

const (
	UserStatusActive  UserStatus = "active"
	UserStatusBlocked UserStatus = "blocked"
	UserStatusPending UserStatus = "pending"
)

type User struct {
	ID           UserID     `json:"id"`
	TenantID     TenantID   `json:"tenant_id"`
	Email        string     `json:"email"`
	Username     string     `json:"username"`
	Phone        string     `json:"phone"`
	PasswordHash string     `json:"-"`
	Roles        []string   `json:"roles"`
	Status       UserStatus `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (u *User) IsActive() bool { return u.Status == UserStatusActive }

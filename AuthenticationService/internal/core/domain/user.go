package domain

import "time"

type UserID string

type User struct {
	ID           UserID
	Tenant       TenantID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	Disabled     bool
}

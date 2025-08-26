package domain

import "time"

type RefreshTokenID string

type RefreshToken struct {
	ID        RefreshTokenID `json:"id"`
	TenantID  TenantID       `json:"tenant_id"`
	UserID    UserID         `json:"user_id"`
	ClientID  ClientID       `json:"client_id"`
	Scopes    []string       `json:"scopes"`
	ExpiresAt time.Time      `json:"expires_at"`
	Revoked   bool           `json:"revoked"`
	CreatedAt time.Time      `json:"created_at"`
}

type AccessToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

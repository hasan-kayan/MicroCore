package domain

import "time"

type ClientID string

type ClientType string

const (
	ClientTypePublic       ClientType = "public"
	ClientTypeConfidential ClientType = "confidential"
)

type Client struct {
	ID            ClientID      `json:"id"`
	TenantID      TenantID      `json:"tenant_id"`
	Name          string        `json:"name"`
	Type          ClientType    `json:"type"`
	SecretHash    string        `json:"-"`
	RedirectURIs  []string      `json:"redirect_uris"`
	AllowedScopes []string      `json:"allowed_scopes"`
	AccessTTL     time.Duration `json:"access_ttl"`
	RefreshTTL    time.Duration `json:"refresh_ttl"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func (c *Client) IsConfidential() bool { return c.Type == ClientTypeConfidential }

package domain

import (
	"time"
)

type TenantID string

type Tenant struct {
	ID        TenantID  `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Tenant) Activate()   { t.Active = true }
func (t *Tenant) Deactivate() { t.Active = false }

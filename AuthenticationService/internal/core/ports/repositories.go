package ports

import (
	"context"

	"github.com/hasan-kayan/MicroCore/AuthenticationService/internal/core/domain"
)

type TenantRepository interface {
	GetByID(ctx context.Context, id domain.TenantID) (*domain.Tenant, error)
	GetBySlug(ctx context.Context, slug string) (*domain.Tenant, error)
	Save(ctx context.Context, t *domain.Tenant) error
}

type ClientRepository interface {
	GetByID(ctx context.Context, id domain.ClientID) (*domain.Client, error)
	Save(ctx context.Context, c *domain.Client) error
}

type UserRepository interface {
	GetByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	GetByEmail(ctx context.Context, tenantID domain.TenantID, email string) (*domain.User, error)
	Save(ctx context.Context, u *domain.User) error
}

type RefreshTokenRepository interface {
	Save(ctx context.Context, rt *domain.RefreshToken) error
	Get(ctx context.Context, id domain.RefreshTokenID) (*domain.RefreshToken, error)
	Revoke(ctx context.Context, id domain.RefreshTokenID) error
}

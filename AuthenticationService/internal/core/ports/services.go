package ports

import (
	"context"
	"time"

	"github.com/hasan-kayan/MicroCore/AuthenticationService/internal/core/domain"
)

type PasswordHasher interface {
	Hash(plain string) (string, error)
	Compare(hash, plain string) error
}

type TokenSigner interface {
	SignAccess(ctx context.Context, claims map[string]any, ttl time.Duration) (token string, exp time.Time, err error)
}

type IDProvider interface {
	NewUserID() domain.UserID
	NewTenantID() domain.TenantID
	NewClientID() domain.ClientID
	NewRefreshTokenID() domain.RefreshTokenID
}

type TimeProvider interface {
	Now() time.Time
}

type TxManager interface {
	InTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Logger interface {
	Info(msg string, kv ...any)
	Error(msg string, kv ...any)
}

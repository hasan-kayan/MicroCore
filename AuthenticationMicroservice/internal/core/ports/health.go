package ports

import (
	"context"

	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/domain"
)

type HealthRepository interface {
	Insert(ctx context.Context, event *domain.HealthEvent) (string, error)
}

type RequestContext interface {
	RemoteIP() string
	RequestID() string
	UserAgent() string
}

type HealthService interface {
	Record(ctx context.Context, payload domain.HealthPayload, info RequestContext) (string, error)
}

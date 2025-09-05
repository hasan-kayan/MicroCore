package services

import (
	"context"
	"errors"
	"time"

	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/domain"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/ports"
)

var ErrInvalidPayload = errors.New("invalid payload")

type healthService struct {
	repo ports.HealthRepository
}

func NewHealthService(r ports.HealthRepository) ports.HealthService {
	return &healthService{repo: r}
}

func (s *healthService) Record(ctx context.Context, payload domain.HealthPayload, info ports.RequestContext) (string, error) {
	if payload.ClientID == "" || payload.Platform == "" || payload.Status == "" {
		return "", ErrInvalidPayload
	}
	ev := domain.HealthEvent{
		ReceivedAt: time.Now().UTC(),
		RemoteIP:   info.RemoteIP,
		RequestID:  info.RequestID,
		UserAgent:  info.UserAgent,
		Payload:    payload,
	}
	return s.repo.Insert(ctx, ev)
}

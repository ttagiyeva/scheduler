package repository

import (
	"context"

	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
)

type Scheduler interface {
	Save(ctx context.Context, s *domain.Scheduler) error
	Get(ctx context.Context, orderName string) (*domain.Scheduler, error)
	GetAll(ctx context.Context, path string, op string, value interface{}) ([]*domain.Scheduler, error)
	Update(ctx context.Context, s *domain.Scheduler) error
	Delete(ctx context.Context, orderName string) error
}

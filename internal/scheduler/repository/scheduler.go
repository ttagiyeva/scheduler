package repository

import (
	"context"

	"github.com/ttagiyeva/scheduler/internal/scheduler/domain"
)

//Scheduler is an interface for scheduler repository
type Scheduler interface {
	Save(ctx context.Context, s *domain.Scheduler) error
	Get(ctx context.Context, orderName string) (*domain.Scheduler, error)
	GetQueried(ctx context.Context, path string, op string, value interface{}) ([]*domain.Scheduler, error)
	Update(ctx context.Context, s *domain.Scheduler) error
	Delete(ctx context.Context, orderName string) error
}

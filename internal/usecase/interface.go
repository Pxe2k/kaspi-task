package usecase

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

type repository interface {
	Find(ctx context.Context, name string) ([]domain.Person, error)
	Get(ctx context.Context, docID string) (domain.Person, error)
	Store(ctx context.Context, person domain.Person) error
}

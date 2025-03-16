package delivery

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

type UseCase interface {
	Get(ctx context.Context, docID string) (domain.Person, error)
	Store(ctx context.Context, req StoreRequest) error
	Find(ctx context.Context, name string) ([]domain.Person, error)
	Check(ctx context.Context, docID string) (string, string, error)
}

package person

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

func (u UseCase) Find(ctx context.Context, name string) ([]domain.Person, error) {
	return u.repository.Find(ctx, name)
}

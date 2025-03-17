package person

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

func (u UseCase) Get(ctx context.Context, docID string) (domain.Person, error) {
	return u.repository.Get(ctx, docID)
}

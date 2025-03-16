package usecase

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/delivery"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

func (u UseCase) Store(ctx context.Context, req delivery.StoreRequest) error {
	return u.repository.Store(ctx, domain.Person{
		DocumentID: req.DocumentID,
		Name:       req.Name,
		Phone:      req.Phone,
	})
}

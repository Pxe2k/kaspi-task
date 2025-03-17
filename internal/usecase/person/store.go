package person

import (
	"context"
	"github.com/Pxe2k/kaspi-task/internal/delivery/person"
	"github.com/Pxe2k/kaspi-task/internal/domain"
)

func (u UseCase) Store(ctx context.Context, req person.StoreRequest) error {
	return u.repository.Store(ctx, domain.Person{
		DocumentID: req.DocumentID,
		Name:       req.Name,
		Phone:      req.Phone,
	})
}

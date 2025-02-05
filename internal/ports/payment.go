package ports

import (
	"context"
	domain "hexagonal/internal/application/core/domain/models"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}

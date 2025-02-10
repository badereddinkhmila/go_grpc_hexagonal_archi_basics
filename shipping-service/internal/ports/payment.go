package ports

import (
	"context"
	domain "shipping-service/internal/application/core/domain/models"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}

package ports

import (
	"context"
	domain "payment-service/internal/application/core/domain/models"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}

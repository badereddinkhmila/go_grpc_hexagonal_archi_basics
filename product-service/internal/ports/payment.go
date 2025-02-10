package ports

import (
	"context"
	domain "product-service/internal/application/core/domain/model"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Product) error
}

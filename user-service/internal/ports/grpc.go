package ports

import (
	"context"
	domain "user-service/internal/application/core/domain/models"
)

type GrpcPort interface {
	GetOrders(ctx context.Context) ([]*domain.Order, error)
}

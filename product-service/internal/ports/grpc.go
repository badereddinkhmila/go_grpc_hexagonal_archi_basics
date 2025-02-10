package ports

import (
	"context"
	domain "product-service/internal/application/core/domain/model"
)

type GrpcPort interface {
	GetProducts(ctx context.Context) ([]*domain.Product, error)
	GetCategories(ctx context.Context) ([]*domain.Category, error)
}

package ports

import (
	"context"
	db "product-service/internal/adapters/db"
	domain "product-service/internal/application/core/domain/model"
)

type PostgreSQLPort interface {
	CreateProduct(ctx context.Context, productInput *domain.Product) (*db.Product, error)
	GetProducts(ctx context.Context) ([]*db.Product, error)
	GetProduct(ctx context.Context, id string) (*db.Product, error)
	UpdateProduct(ctx context.Context, product *domain.Product) (*db.Product, error)
	DeleteProduct(ctx context.Context, id string) (bool, error)
	CreateCategory(ctx context.Context, categoryInput *domain.Category) (*db.Category, error)
	GetCategories(ctx context.Context) ([]*db.Category, error)
	GetCategory(ctx context.Context, id string) (*db.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) (*db.Category, error)
	DeleteCategory(ctx context.Context, id string) (bool, error)
}

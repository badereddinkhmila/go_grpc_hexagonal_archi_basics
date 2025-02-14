package db

import (
	"context"
	domain "product-service/internal/application/core/domain/model"
)

func (a *PostgreSQLAdapter) CreateProduct(ctx context.Context, productInput *domain.Product) (*Product, error) {
	return &Product{}, nil
}

func (a *PostgreSQLAdapter) GetProducts(ctx context.Context) ([]*Product, error) {
	return []*Product{}, nil
}

func (a *PostgreSQLAdapter) GetProduct(ctx context.Context, id string) (*Product, error) {
	return &Product{}, nil
}

func (a *PostgreSQLAdapter) UpdateProduct(ctx context.Context, product *domain.Product) (*Product, error) {
	return &Product{}, nil
}

func (a *PostgreSQLAdapter) DeleteProduct(ctx context.Context, id string) (bool, error) {
	return true, nil
}

package db

import (
	"context"
	domain "product-service/internal/application/core/domain/model"
)

func (a *PostgreSQLAdapter) CreateCategory(ctx context.Context, categoryInput *domain.Category) (*Category, error) {
	return &Category{}, nil
}

func (a *PostgreSQLAdapter) GetCategories(ctx context.Context) ([]*Category, error) {
	return []*Category{}, nil
}

func (a *PostgreSQLAdapter) GetCategory(ctx context.Context, id string) (*Category, error) {
	return &Category{}, nil
}

func (a *PostgreSQLAdapter) UpdateCategory(ctx context.Context, category *domain.Category) (*Category, error) {
	return &Category{}, nil
}

func (a *PostgreSQLAdapter) DeleteCategory(ctx context.Context, id string) (bool, error) {
	return true, nil
}

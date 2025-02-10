package api

import (
	"context"
	domain "product-service/internal/application/core/domain/model"
	"product-service/internal/ports"
)

type Application struct {
	db ports.PostgreSQLPort
}

func NewApplication(db ports.PostgreSQLPort) *Application {
	return &Application{
		db: db,
	}
}

func (app *Application) GetProducts(ctx context.Context) ([]*domain.Product, error) {
	return nil, nil
}

func (app *Application) GetCategories(ctx context.Context) ([]*domain.Category, error) {
	return nil, nil

}

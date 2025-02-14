package api

import (
	"context"
	"fmt"
	"log"
	domain "product-service/internal/application/core/domain/model"
	dto "product-service/internal/application/core/dto"
	ports "product-service/internal/ports"
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
	products, err := app.db.GetProducts(ctx)
	if err != nil {
		log.Printf("API.GetProducts: error getting products %v", err)
		return nil, err
	}
	productsDomain := make([]*domain.Product, 0)
	for _, product := range products {
		productsDomain = append(productsDomain, &domain.Product{
			ID:    &product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsDomain, nil
}

func (app *Application) GetCategories(ctx context.Context) ([]*domain.Category, error) {
	categories, err := app.db.GetCategories(ctx)
	if err != nil {
		log.Printf("API.GetCategories: error getting categories %v", err)
		return nil, err
	}
	categoriesDomain := make([]*domain.Category, 0)
	for _, category := range categories {
		categoriesDomain = append(categoriesDomain, &domain.Category{
			ID:        category.ID,
			OwnerID:   category.OwnerID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
		})
	}
	return categoriesDomain, nil
}

func (app *Application) CreateProduct(ctx context.Context, createProductDTO dto.CreateProductDTO) ([]*domain.Product, error) {
	fmt.Printf("API.CreateProduct: dto body %v", createProductDTO)
	return nil, nil
}

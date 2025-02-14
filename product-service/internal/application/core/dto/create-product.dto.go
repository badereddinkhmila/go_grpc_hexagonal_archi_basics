package dto

import (
	"log"

	validate "product-service/config/validator"
	domain "product-service/internal/application/core/domain/model"
)

type CreateProductDTO struct {
	OwnerID  uint    `json:"owner_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Quantity int64   `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type UpdateProductDTO struct {
}

func (cp *CreateProductDTO) toDomainProduct() (*domain.Product, error) {
	return domain.NewProduct(cp.OwnerID, cp.Quantity, cp.Name, cp.Price), nil
}

func (cp *CreateProductDTO) Validate() error {
	// Validate struct
	if err := validate.ValidateStruct(&cp); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

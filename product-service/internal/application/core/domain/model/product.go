package domain

import (
	"time"
)

type Product struct {
	ID        *uint      `json:"id"`
	OwnerID   uint       `json:"owner_id" validate:"required"`
	Name      string     `json:"name"`
	Quantity  int64      `json:"quantity" validate:"required"`
	Price     float64    `json:"price" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewProduct(ownerId uint, quantity int64, name string, price float64) *Product {
	return &Product{
		CreatedAt: time.Now(),
		Name:      name,
		OwnerID:   ownerId,
		Quantity:  quantity,
		Price:     price,
		ID:        nil,
	}
}

package domain

import "time"

type Product struct {
	ID        int64   `json:"id"`
	OwnerID   int64   `json:"owner_id" validate:"required"`
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt int64   `json:"created_at"`
}

func NewProduct(ownerId, quantity int64, name string, price float64) Product {
	return Product{
		CreatedAt: time.Now().Unix(),
		Name:      name,
		OwnerID:   ownerId,
		Quantity:  quantity,
		Price:     price,
	}
}

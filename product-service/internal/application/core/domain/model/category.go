package domain

import "time"

type Category struct {
	ID        int64  `json:"id"`
	OwnerID   int64  `json:"owner_id" validate:"required"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}

func NewCategory(ownerId int64, name string) Category {
	return Category{
		CreatedAt: time.Now().Unix(),
		Name:      "Pending",
		OwnerID:   ownerId,
	}
}

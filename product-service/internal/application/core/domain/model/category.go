package domain

import "time"

type Category struct {
	ID        uint      `json:"id"`
	OwnerID   uint      `json:"owner_id" validate:"required"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategory(ownerId uint, name string) Category {
	return Category{
		CreatedAt: time.Now(),
		Name:      "Pending",
		OwnerID:   ownerId,
	}
}

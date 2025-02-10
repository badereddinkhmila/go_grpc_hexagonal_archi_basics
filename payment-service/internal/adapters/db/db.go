package db

import (
	"context"
	"fmt"
	"payment-service/config"
	domain "payment-service/internal/application/core/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	Items      []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductID string
	UnitPrice float32
	Quantity  int64
	OrderID   uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter() (*Adapter, error) {
	c := config.PostgreSQL()

	config := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, "disable",
	)
	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	err = database.AutoMigrate(OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: database}, nil
}

func (a *Adapter) CreateOrder(id string) (string, error) {
	return "", nil
}
func (a *Adapter) UpdateOrder(id string) (string, error) {
	return "", nil
}

func (a *Adapter) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	return []*domain.Order{}, nil
}

package db

import (
	"context"
	"fmt"
	config "product-service/config"
	domain "product-service/internal/application/core/domain/model"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID        pgtype.UUID `gorm:"primarykey"`
	OwnerID   int64
	Name      string
	Quantity  int64
	Price     decimal.Decimal
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        pgtype.UUID `gorm:"primarykey"`
	OwnerID   string
	Name      string `gorm:"index"`
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PostgreSQLAdapter struct {
	db *gorm.DB
}

func NewPostgreSQLAdapter() (*PostgreSQLAdapter, error) {
	c := config.PostgreSQL()

	config := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, "disable",
	)
	database, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	err = database.AutoMigrate(Product{}, Category{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &PostgreSQLAdapter{db: database}, nil
}

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

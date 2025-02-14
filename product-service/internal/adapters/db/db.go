package db

import (
	"fmt"
	config "product-service/config"

	pq "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	OwnerID  uint
	Name     string `gorm:"index"`
	Quantity int64
	Price    float64

	Categories []Category `gorm:"many2many:categories_products;"`
}

type Category struct {
	*gorm.Model
	OwnerID  uint
	Name     string         `gorm:"index"`
	Tags     pq.StringArray `gorm:"type:text[]"`
	Products []Product      `gorm:"many2many:categories_products;"`
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

	err = database.AutoMigrate(&Product{}, &Category{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &PostgreSQLAdapter{db: database}, nil
}

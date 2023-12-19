package repository

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	entities "example.com/products/internal/domain/entities"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, name string, price float64) (int, error)
	GetProducts(ctx context.Context) ([]entities.Product, error)
}

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *sql.DB) ProductRepository {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	gormDB.AutoMigrate(&entities.Product{})

	return &gormProductRepository{db: gormDB}
}

func (r *gormProductRepository) CreateProduct(ctx context.Context, name string, price float64) (int, error) {
	product := entities.Product{Name: name, Price: price}
	result := r.db.WithContext(ctx).Create(&product)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(product.ID), nil
}

func (r *gormProductRepository) GetProducts(ctx context.Context) ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.WithContext(ctx).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

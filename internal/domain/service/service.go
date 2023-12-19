package service

import (
	"context"

	"example.com/products/internal/app/router/dto"
	entities "example.com/products/internal/domain/entities"
	"example.com/products/internal/infrastructure/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, dto *dto.CreateProductDTO) (int, error)
	GetProducts(ctx context.Context) ([]entities.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, dto *dto.CreateProductDTO) (int, error) {
	return s.repo.CreateProduct(ctx, dto.Name, dto.Price)
}

func (s *productService) GetProducts(ctx context.Context) ([]entities.Product, error) {
	return s.repo.GetProducts(ctx)
}

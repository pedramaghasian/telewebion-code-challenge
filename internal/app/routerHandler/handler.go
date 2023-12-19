package handler

import (
	"context"
	"net/http"

	"example.com/products/internal/app/router/dto"
	"example.com/products/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context, dto *dto.CreateProductDTO) {
	productID, err := h.service.CreateProduct(context.Background(), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": productID})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetProducts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

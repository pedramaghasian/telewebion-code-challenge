package router

import (
	"fmt"
	"net/http"

	"example.com/products/internal/app/router/dto"
	handler "example.com/products/internal/app/routerHandler"
	"example.com/products/pkg/rateLimiter"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine         *gin.Engine
	createCounter  *rateLimiter.Counter
	readCounter    *rateLimiter.Counter
	productHandler *handler.ProductHandler
}

func NewRouter(createCounter, readCounter *rateLimiter.Counter, productHandler *handler.ProductHandler) *Router {
	return &Router{
		engine:         gin.Default(),
		createCounter:  createCounter,
		readCounter:    readCounter,
		productHandler: productHandler,
	}
}

func (r *Router) SetupRoutes() {
	r.engine.POST("/products", r.handleCreateProduct)
	r.engine.GET("/products", r.handleGetProducts)
}

func (r *Router) Run(port int) {
	r.engine.Run(fmt.Sprintf(":%d", port))
}

func (r *Router) handleCreateProduct(c *gin.Context) {
	if r.createCounter.IsLimitExceeded() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		return
	}
	r.createCounter.Increment()

	var request dto.CreateProductDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r.productHandler.CreateProduct(c, &request)
}

func (r *Router) handleGetProducts(c *gin.Context) {
	if r.readCounter.IsLimitExceeded() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		return
	}
	r.readCounter.Increment()

	r.productHandler.GetProducts(c)
}

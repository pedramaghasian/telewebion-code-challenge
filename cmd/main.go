package main

import (
	"database/sql"
	"fmt"

	"example.com/products/config"
	"example.com/products/internal/app/router"
	handler "example.com/products/internal/app/routerHandler"
	"example.com/products/internal/domain/service"
	"example.com/products/internal/infrastructure/repository"
	"example.com/products/pkg/rateLimiter"
	_ "github.com/lib/pq"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Create a PostgreSQL connection
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.User,
		appConfig.Database.Password,
		appConfig.Database.DBName,
		appConfig.Database.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	productRepo := repository.NewGormProductRepository(db)

	createCounter := rateLimiter.NewCounter(50)
	readCounter := rateLimiter.NewCounter(50)

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService)

	productRouter := router.NewRouter(createCounter, readCounter, productHandler)

	productRouter.SetupRoutes()

	productRouter.Run(appConfig.Server.Port)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	httpAdapter "kasir-api/internal/adapter/http"
	"kasir-api/internal/infrastructure/config"
	"kasir-api/internal/infrastructure/persistence/postgres"
	categoryUseCase "kasir-api/internal/usecase/category"
	productUseCase "kasir-api/internal/usecase/product"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := postgres.InitDB(cfg.DbConn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories (infrastructure layer)
	productRepo := postgres.NewProductRepository(db)
	categoryRepo := postgres.NewCategoryRepository(db)

	// Initialize use cases (application layer)
	// Repositories implement the interfaces defined in domain layer
	productService := productUseCase.NewService(productRepo)
	categoryService := categoryUseCase.NewService(categoryRepo)

	// Initialize HTTP router (adapter layer)
	router := httpAdapter.NewRouter(productService, categoryService)
	router.SetupRoutes()

	// Start server
	fmt.Println("Starting server on :" + cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

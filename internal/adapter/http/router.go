package http

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/adapter/http/handler"
	"kasir-api/internal/usecase/category"
	"kasir-api/internal/usecase/product"
)

// Router sets up all HTTP routes
type Router struct {
	productHandler  *handler.ProductHandler
	categoryHandler *handler.CategoryHandler
}

// NewRouter creates a new router with handlers
func NewRouter(productUseCase product.ProductUseCase, categoryUseCase category.CategoryUseCase) *Router {
	return &Router{
		productHandler:  handler.NewProductHandler(productUseCase),
		categoryHandler: handler.NewCategoryHandler(categoryUseCase),
	}
}

// SetupRoutes registers all routes with the default mux
func (router *Router) SetupRoutes() {
	// Product routes
	http.HandleFunc("/api/products", router.productHandler.HandleProducts)
	http.HandleFunc("/api/product/", router.productHandler.HandleProductByID)

	// Category routes
	http.HandleFunc("/api/categories", router.categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", router.categoryHandler.HandleCategoryByID)

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})
}

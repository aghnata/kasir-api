package routes

import (
	"kasir-api/handlers"
	"kasir-api/storage"
	"net/http"
)

// SetupRoutes configures all the API routes
func SetupRoutes() {
	// Initialize storage
	produkStorage := storage.NewProdukStorage()

	// Initialize handlers
	produkHandler := handlers.NewProdukHandler(produkStorage)
	healthHandler := handlers.NewHealthHandler()

	// Product routes with specific ID
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			produkHandler.GetProdukByID(w, r)
		case "PUT":
			produkHandler.UpdateProdukByID(w, r)
		case "DELETE":
			produkHandler.DeleteProdukByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Product routes for collection
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			produkHandler.GetAllProduk(w, r)
		case "POST":
			produkHandler.CreateProduk(w, r)
		default:
			http.Error(w, "Method not allowed boy", http.StatusMethodNotAllowed)
		}
	})

	// Health check route
	http.HandleFunc("/health", healthHandler.Health)
}

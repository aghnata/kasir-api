package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/storage"
	"net/http"
	"strconv"
	"strings"
)

// ProdukHandler handles HTTP requests for product operations
type ProdukHandler struct {
	storage *storage.ProdukStorage
}

// NewProdukHandler creates a new ProdukHandler instance
func NewProdukHandler(storage *storage.ProdukStorage) *ProdukHandler {
	return &ProdukHandler{storage: storage}
}

// GetAllProduk handles GET /api/produk
func (ph *ProdukHandler) GetAllProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ph.storage.GetAll())
}

// CreateProduk handles POST /api/produk
func (ph *ProdukHandler) CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru models.Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduk := ph.storage.Create(produkBaru)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdProduk)
}

// GetProdukByID handles GET /api/produk/{id}
func (ph *ProdukHandler) GetProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := ph.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	produk, err := ph.storage.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

// UpdateProdukByID handles PUT /api/produk/{id}
func (ph *ProdukHandler) UpdateProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := ph.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var produkUpdate models.Produk
	err = json.NewDecoder(r.Body).Decode(&produkUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProduk, err := ph.storage.Update(id, produkUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduk)
}

// DeleteProdukByID handles DELETE /api/produk/{id}
func (ph *ProdukHandler) DeleteProdukByID(w http.ResponseWriter, r *http.Request) {
	id, err := ph.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	err = ph.storage.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk berhasil dihapus"})
}

// extractIDFromPath extracts the ID from the URL path
func (ph *ProdukHandler) extractIDFromPath(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/produk/")
	return strconv.Atoi(idStr)
}

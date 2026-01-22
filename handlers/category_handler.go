package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/storage"
	"net/http"
	"strconv"
	"strings"
)

// CategoryHandler handles HTTP requests for category operations
type CategoryHandler struct {
	storage *storage.CategoryStorage
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewCategoryHandler(storage *storage.CategoryStorage) *CategoryHandler {
	return &CategoryHandler{storage: storage}
}

// GetAllCategory handles GET /api/category
func (ch *CategoryHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ch.storage.GetAll())
}

// CreateCategory handles POST /api/category
func (ch *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryBaru models.Category
	err := json.NewDecoder(r.Body).Decode(&categoryBaru)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdCategory := ch.storage.Create(categoryBaru)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCategory)
}

// GetCategoryByID handles GET /api/category/{id}
func (ch *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := ch.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	category, err := ch.storage.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

// UpdateCategoryByID handles PUT /api/category/{id}
func (ch *CategoryHandler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := ch.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var categoryUpdate models.Category
	err = json.NewDecoder(r.Body).Decode(&categoryUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedCategory, err := ch.storage.Update(id, categoryUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCategory)
}

// DeleteCategoryByID handles DELETE /api/category/{id}
func (ch *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := ch.extractIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	err = ch.storage.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Kategori berhasil dihapus"})
}

// extractIDFromPath extracts the ID from the URL path
func (ch *CategoryHandler) extractIDFromPath(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/api/categories/")
	return strconv.Atoi(idStr)
}

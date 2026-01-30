package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/internal/adapter/http/dto"
	"kasir-api/internal/usecase/product"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	useCase product.ProductUseCase
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(useCase product.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

// HandleProducts handles GET /api/products and POST /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleProductByID handles GET/PUT/DELETE /api/product/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetAll handles GET /api/products
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.useCase.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	// Convert to HTTP response DTOs
	response := make([]dto.ProductResponse, len(products))
	for i, p := range products {
		response[i] = dto.ProductResponse{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
			Stock: p.Stock,
		}
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Create handles POST /api/products
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Convert HTTP DTO to usecase input
	input := product.CreateProductInput{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	output, err := h.useCase.Create(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.ProductResponse{
		ID:    output.ID,
		Name:  output.Name,
		Price: output.Price,
		Stock: output.Stock,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// GetByID handles GET /api/product/{id}
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/product/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	output, err := h.useCase.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	response := dto.ProductResponse{
		ID:    output.ID,
		Name:  output.Name,
		Price: output.Price,
		Stock: output.Stock,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Update handles PUT /api/product/{id}
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/product/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	input := product.UpdateProductInput{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	output, err := h.useCase.Update(id, input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.ProductResponse{
		ID:    output.ID,
		Name:  output.Name,
		Price: output.Price,
		Stock: output.Stock,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Delete handles DELETE /api/product/{id}
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/product/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dto.MessageResponse{
		Message: "Product deleted successfully",
	})
}

// Helper functions
func extractIDFromPath(path, prefix string) (int, error) {
	idStr := strings.TrimPrefix(path, prefix)
	return strconv.Atoi(idStr)
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, dto.ErrorResponse{Error: message})
}

package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/adapter/http/dto"
	"kasir-api/internal/usecase/category"
)

// CategoryHandler handles HTTP requests for categories
type CategoryHandler struct {
	useCase category.CategoryUseCase
}

// NewCategoryHandler creates a new CategoryHandler
func NewCategoryHandler(useCase category.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{useCase: useCase}
}

// HandleCategories handles GET /api/categories and POST /api/categories
func (h *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleCategoryByID handles GET/PUT/DELETE /api/category/{id}
func (h *CategoryHandler) HandleCategoryByID(w http.ResponseWriter, r *http.Request) {
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

// GetAll handles GET /api/categories
func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.useCase.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}

	// Convert to HTTP response DTOs
	response := make([]dto.CategoryResponse, len(categories))
	for i, c := range categories {
		response[i] = dto.CategoryResponse{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		}
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Create handles POST /api/categories
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Convert HTTP DTO to usecase input
	input := category.CreateCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	}

	output, err := h.useCase.Create(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.CategoryResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// GetByID handles GET /api/category/{id}
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/category/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	output, err := h.useCase.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	response := dto.CategoryResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Update handles PUT /api/category/{id}
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/category/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var req dto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	input := category.UpdateCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	}

	output, err := h.useCase.Update(id, input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.CategoryResponse{
		ID:          output.ID,
		Name:        output.Name,
		Description: output.Description,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Delete handles DELETE /api/category/{id}
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromPath(r.URL.Path, "/api/category/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	if err := h.useCase.Delete(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dto.MessageResponse{
		Message: "Category deleted successfully",
	})
}

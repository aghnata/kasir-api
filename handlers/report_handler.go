package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type ReportHandler struct {
	services *services.ReportService
}

func NewReportHandler(services *services.ReportService) *ReportHandler {
	return &ReportHandler{services: services}
}

func (h *ReportHandler) HandleGetReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	report, err := h.services.GenerateSalesReport()
	if err != nil {
		http.Error(w, "Failed to generate report", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

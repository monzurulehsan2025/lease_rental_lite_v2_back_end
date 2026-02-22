package handler

import (
	"encoding/json"
	"net/http"

	"github.com/securerent/risk-assessment-service/internal/model"
	"github.com/securerent/risk-assessment-service/internal/service"
)

type RiskHandler struct {
	riskService service.RiskService
}

func NewRiskHandler(rs service.RiskService) *RiskHandler {
	return &RiskHandler{riskService: rs}
}

func (h *RiskHandler) Assess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var profile model.RenterProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.riskService.AssessRenter(profile)
	if err != nil {
		http.Error(w, "Failed to assess risk", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *RiskHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

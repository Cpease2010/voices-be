package handlers

import (
	"encoding/json"
	"net/http"
	"voices/db"
	"voices/services"
)

type EngagementHandler struct {
	Service services.EngagementService
}

func NewEngagementHandler(service services.EngagementService) *EngagementHandler {
	return &EngagementHandler{Service: service}
}

func (h *EngagementHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	engagements, err := h.Service.GetAllEngagements()
	if err != nil {
		http.Error(w, "Failed to retrieve engagements", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(engagements)
}

func (h *EngagementHandler) Create(w http.ResponseWriter, r *http.Request) {
	var e db.Engagement
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateEngagement(e); err != nil {
		http.Error(w, "Failed to create engagement", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

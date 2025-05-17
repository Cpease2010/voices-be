package handlers

import (
	"encoding/json"
	"net/http"
	"voices/db"
	"voices/services"
)

type TrusteeHandler struct {
	Service services.TrusteeService
}

func NewTrusteeHandler(service services.TrusteeService) *TrusteeHandler {
	return &TrusteeHandler{Service: service}
}

func (h *TrusteeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	trustees, err := h.Service.GetAllTrustees()
	if err != nil {
		http.Error(w, "Failed to fetch trustees", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(trustees)
}

func (h *TrusteeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var t db.Trustee
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateTrustee(t); err != nil {
		http.Error(w, "Failed to create trustee", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

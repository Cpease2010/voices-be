package handlers

import (
	"encoding/json"
	"net/http"
	"voices/db"
	"voices/services"
)

type CitizenHandler struct {
	Service services.CitizenService
}

func NewCitizenHandler(service services.CitizenService) *CitizenHandler {
	return &CitizenHandler{Service: service}
}

func (h *CitizenHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	citizens, err := h.Service.GetAllCitizens()
	if err != nil {
		http.Error(w, "Failed to retrieve citizens", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(citizens)
}

func (h *CitizenHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c db.Citizen
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateCitizen(c); err != nil {
		http.Error(w, "Failed to create citizen", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

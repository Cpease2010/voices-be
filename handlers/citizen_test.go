package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"voices/db"
)

type mockCitizenService struct {
	GetAllFunc func() ([]db.Citizen, error)
	CreateFunc func(c db.Citizen) error
}

func (m *mockCitizenService) GetAllCitizens() ([]db.Citizen, error) {
	return m.GetAllFunc()
}

func (m *mockCitizenService) CreateCitizen(c db.Citizen) error {
	return m.CreateFunc(c)
}

func TestGetAllCitizens(t *testing.T) {
	mockService := &mockCitizenService{
		GetAllFunc: func() ([]db.Citizen, error) {
			return []db.Citizen{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
			}, nil
		},
	}
	handler := NewCitizenHandler(mockService)

	req := httptest.NewRequest("GET", "/citizens", nil)
	rec := httptest.NewRecorder()

	handler.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rec.Code)
	}
	var citizens []db.Citizen
	if err := json.NewDecoder(rec.Body).Decode(&citizens); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(citizens) != 2 {
		t.Errorf("expected 2 citizens, got %d", len(citizens))
	}
}

func TestCreateCitizen(t *testing.T) {
	mockService := &mockCitizenService{
		CreateFunc: func(c db.Citizen) error {
			if c.Name == "" {
				return errors.New("name required")
			}
			return nil
		},
	}
	handler := NewCitizenHandler(mockService)

	citizen := db.Citizen{Name: "Alice"}
	body, _ := json.Marshal(citizen)
	req := httptest.NewRequest("POST", "/citizens/create", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rec.Code)
	}
}

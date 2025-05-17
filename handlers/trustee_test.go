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

type mockTrusteeService struct {
	GetAllFunc func() ([]db.Trustee, error)
	CreateFunc func(t db.Trustee) error
}

func (m *mockTrusteeService) GetAllTrustees() ([]db.Trustee, error) {
	return m.GetAllFunc()
}

func (m *mockTrusteeService) CreateTrustee(t db.Trustee) error {
	return m.CreateFunc(t)
}

func TestGetAllTrustees(t *testing.T) {
	mockService := &mockTrusteeService{
		GetAllFunc: func() ([]db.Trustee, error) {
			return []db.Trustee{
				{ID: 1, Name: "Trustee A"},
				{ID: 2, Name: "Trustee B"},
			}, nil
		},
	}

	handler := NewTrusteeHandler(mockService)

	req := httptest.NewRequest("GET", "/trustees", nil)
	rec := httptest.NewRecorder()

	handler.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rec.Code)
	}

	var trustees []db.Trustee
	if err := json.NewDecoder(rec.Body).Decode(&trustees); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(trustees) != 2 {
		t.Errorf("expected 2 trustees, got %d", len(trustees))
	}
}

func TestCreateTrustee(t *testing.T) {
	mockService := &mockTrusteeService{
		CreateFunc: func(t db.Trustee) error {
			if t.Name == "" {
				return errors.New("name required")
			}
			return nil
		},
	}

	handler := NewTrusteeHandler(mockService)

	payload := db.Trustee{Name: "New Trustee"}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/trustees/create", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rec.Code)
	}
}

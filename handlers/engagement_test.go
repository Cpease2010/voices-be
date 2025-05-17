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

type mockEngagementService struct {
	GetAllFunc func() ([]db.Engagement, error)
	CreateFunc func(e db.Engagement) error
}

func (m *mockEngagementService) GetAllEngagements() ([]db.Engagement, error) {
	return m.GetAllFunc()
}

func (m *mockEngagementService) CreateEngagement(e db.Engagement) error {
	return m.CreateFunc(e)
}

func TestGetAllEngagements(t *testing.T) {
	mockService := &mockEngagementService{
		GetAllFunc: func() ([]db.Engagement, error) {
			return []db.Engagement{
				{ID: 1, TrusteeID: 10, CitizenID: 20, Feedback: "Great"},
				{ID: 2, TrusteeID: 11, CitizenID: 21, Feedback: "Okay"},
			}, nil
		},
	}

	handler := NewEngagementHandler(mockService)

	req := httptest.NewRequest("GET", "/engagements", nil)
	rec := httptest.NewRecorder()

	handler.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rec.Code)
	}

	var engagements []db.Engagement
	if err := json.NewDecoder(rec.Body).Decode(&engagements); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(engagements) != 2 {
		t.Errorf("expected 2 engagements, got %d", len(engagements))
	}
}

func TestCreateEngagement(t *testing.T) {
	mockService := &mockEngagementService{
		CreateFunc: func(e db.Engagement) error {
			if e.TrusteeID == 0 || e.CitizenID == 0 || e.Feedback == "" {
				return errors.New("invalid input")
			}
			return nil
		},
	}

	handler := NewEngagementHandler(mockService)

	payload := db.Engagement{
		TrusteeID: 10,
		CitizenID: 20,
		Feedback:  "Helpful and responsive",
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest("POST", "/engagements/create", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rec.Code)
	}
}

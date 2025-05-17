package handlers

import (
	"encoding/json"
	"net/http"
	"voices/db"
)

type CreateTrusteeRequest struct {
	UserID       int64  `json:"user_id"` // FK to users.id
	Name         string `json:"name"`
	Position     string `json:"position"`
	WorkLocation string `json:"work_location"`
}

type TrusteeResponse struct {
	ID           int64  `json:"id"`
	UserID       int64  `json:"user_id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	WorkLocation string `json:"work_location"`
	CreatedAt    string `json:"created_at"`
}

func HandleTrustees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetTrustees(w, r)
	case http.MethodPost:
		handleCreateTrustee(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCreateTrustee(w http.ResponseWriter, r *http.Request) {
	var req CreateTrusteeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec(`
        INSERT INTO trustee_profiles (user_id, name, position, work_location)
        VALUES (?, ?, ?, ?)`,
		req.UserID, req.Name, req.Position, req.WorkLocation,
	)
	if err != nil {
		http.Error(w, "Failed to insert trustee", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func handleGetTrustees(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.DB.Query(`
        SELECT id, user_id, name, position, work_location, created_at
        FROM trustee_profiles ORDER BY created_at DESC`)
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var trustees []TrusteeResponse
	for rows.Next() {
		var t TrusteeResponse
		if err := rows.Scan(&t.ID, &t.UserID, &t.Name, &t.Position, &t.WorkLocation, &t.CreatedAt); err != nil {
			http.Error(w, "Row scan failed", http.StatusInternalServerError)
			return
		}
		trustees = append(trustees, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trustees)
}

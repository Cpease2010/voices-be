package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voices/db"
)

type CreateEngagementRequest struct {
	CitizenID int64    `json:"citizen_id"`
	TrusteeID int64    `json:"trustee_id"`
	Category  string   `json:"category"` // "positive", "neutral", "negative"
	Comment   string   `json:"comment"`
	Tags      []string `json:"tags"`
	Location  string   `json:"location"`
}

type EngagementResponse struct {
	ID        int64    `json:"id"`
	CitizenID int64    `json:"citizen_id"`
	TrusteeID int64    `json:"trustee_id"`
	Category  string   `json:"category"`
	Comment   string   `json:"comment"`
	Tags      []string `json:"tags"`
	Location  string   `json:"location"`
	CreatedAt string   `json:"created_at"`
}

func HandleEngagements(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleCreateEngagement(w, r)
	case http.MethodGet:
		handleGetEngagements(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCreateEngagement(w http.ResponseWriter, r *http.Request) {
	var req CreateEngagementRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	tagsJson, err := json.Marshal(req.Tags)
	if err != nil {
		http.Error(w, "Failed to encode tags", http.StatusInternalServerError)
		return
	}

	result, err := db.DB.Exec(`
        INSERT INTO engagements (citizen_id, trustee_id, category, comment, tags, location)
        VALUES (?, ?, ?, ?, ?, ?)`,
		req.CitizenID, req.TrusteeID, req.Category, req.Comment, string(tagsJson), req.Location,
	)
	if err != nil {
		http.Error(w, "Failed to insert engagement", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func handleGetEngagements(w http.ResponseWriter, r *http.Request) {
	trusteeIDStr := r.URL.Query().Get("trustee_id")
	if trusteeIDStr == "" {
		http.Error(w, "trustee_id required", http.StatusBadRequest)
		return
	}

	trusteeID, err := strconv.ParseInt(trusteeIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid trustee_id", http.StatusBadRequest)
		return
	}

	rows, err := db.DB.Query(`
        SELECT id, citizen_id, trustee_id, category, comment, tags, location, created_at
        FROM engagements
        WHERE trustee_id = ? ORDER BY created_at DESC`,
		trusteeID,
	)
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var engagements []EngagementResponse
	for rows.Next() {
		var e EngagementResponse
		var tagsStr string
		if err := rows.Scan(&e.ID, &e.CitizenID, &e.TrusteeID, &e.Category, &e.Comment, &tagsStr, &e.Location, &e.CreatedAt); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(tagsStr), &e.Tags)
		engagements = append(engagements, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(engagements)
}

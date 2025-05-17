package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"voices/db"
)

type CitizenEngagement struct {
	ID        int64    `json:"id"`
	TrusteeID int64    `json:"trustee_id"`
	Category  string   `json:"category"`
	Comment   string   `json:"comment"`
	Tags      []string `json:"tags"`
	Location  string   `json:"location"`
	CreatedAt string   `json:"created_at"`
}

type CitizenProfile struct {
	ID          int64               `json:"id"`
	Engagements []CitizenEngagement `json:"engagements"`
}

func HandleCitizen(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract citizen ID from the URL path
	path := strings.TrimPrefix(r.URL.Path, "/citizens/")
	id, err := strconv.ParseInt(path, 10, 64)
	if err != nil {
		http.Error(w, "Invalid citizen ID", http.StatusBadRequest)
		return
	}

	// Optional: check if the user exists
	var exists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists)
	if err != nil || !exists {
		http.Error(w, "Citizen not found", http.StatusNotFound)
		return
	}

	// Fetch all engagements tied to this citizen
	rows, err := db.DB.Query(`
        SELECT id, trustee_id, category, comment, tags, location, created_at
        FROM engagements
        WHERE citizen_id = ?
        ORDER BY created_at DESC`, id)
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var engagements []CitizenEngagement
	for rows.Next() {
		var e CitizenEngagement
		var tagsStr string
		if err := rows.Scan(&e.ID, &e.TrusteeID, &e.Category, &e.Comment, &tagsStr, &e.Location, &e.CreatedAt); err != nil {
			http.Error(w, "Row scan failed", http.StatusInternalServerError)
			return
		}
		json.Unmarshal([]byte(tagsStr), &e.Tags)
		engagements = append(engagements, e)
	}

	profile := CitizenProfile{
		ID:          id,
		Engagements: engagements,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

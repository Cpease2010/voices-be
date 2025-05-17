package services

import (
	"voices/db"
)

type EngagementService interface {
	GetAllEngagements() ([]db.Engagement, error)
	CreateEngagement(e db.Engagement) error
}

type engagementService struct{}

func NewEngagementService() EngagementService {
	return &engagementService{}
}

func (s *engagementService) GetAllEngagements() ([]db.Engagement, error) {
	rows, err := db.DB.Query("SELECT id, trustee_id, citizen_id, feedback FROM engagements")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var engagements []db.Engagement
	for rows.Next() {
		var e db.Engagement
		if err := rows.Scan(&e.ID, &e.TrusteeID, &e.CitizenID, &e.Feedback); err != nil {
			return nil, err
		}
		engagements = append(engagements, e)
	}
	return engagements, nil
}

func (s *engagementService) CreateEngagement(e db.Engagement) error {
	_, err := db.DB.Exec("INSERT INTO engagements (trustee_id, citizen_id, feedback) VALUES (?, ?, ?)", e.TrusteeID, e.CitizenID, e.Feedback)
	return err
}

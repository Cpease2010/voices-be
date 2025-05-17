package services

import "voices/db"

type TrusteeService interface {
	GetAllTrustees() ([]db.Trustee, error)
	CreateTrustee(t db.Trustee) error
}

type trusteeService struct{}

func NewTrusteeService() TrusteeService {
	return &trusteeService{}
}

func (s *trusteeService) GetAllTrustees() ([]db.Trustee, error) {
	rows, err := db.DB.Query("SELECT * FROM trustees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trustees []db.Trustee
	for rows.Next() {
		var t db.Trustee
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, err
		}
		trustees = append(trustees, t)
	}
	return trustees, nil
}

func (s *trusteeService) CreateTrustee(t db.Trustee) error {
	_, err := db.DB.Exec("INSERT INTO trustees (name) VALUES (?)", t.Name)
	return err
}

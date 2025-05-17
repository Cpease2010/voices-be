package services

import (
	"voices/db"
)

type CitizenService interface {
	GetAllCitizens() ([]db.Citizen, error)
	CreateCitizen(c db.Citizen) error
}

type citizenService struct{}

func NewCitizenService() CitizenService {
	return &citizenService{}
}

func (s *citizenService) GetAllCitizens() ([]db.Citizen, error) {
	rows, err := db.DB.Query("SELECT id, name FROM citizens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var citizens []db.Citizen
	for rows.Next() {
		var c db.Citizen
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		citizens = append(citizens, c)
	}
	return citizens, nil
}

func (s *citizenService) CreateCitizen(c db.Citizen) error {
	_, err := db.DB.Exec("INSERT INTO citizens (name) VALUES (?)", c.Name)
	return err
}

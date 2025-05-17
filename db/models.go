package db

type Trustee struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	WorkLocation string `json:"work_location"`
}

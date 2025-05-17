package db

type Trustee struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Position     string `json:"position"`
	WorkLocation string `json:"work_location"`
}

type Citizen struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Engagement struct {
	ID        int    `json:"id"`
	TrusteeID int    `json:"trustee_id"`
	CitizenID int    `json:"citizen_id"`
	Feedback  string `json:"feedback"`
}

package model

type Dealer struct {
	Id                string `json:"id"`
	DealerName        string `json:"dealer"`
	StreetAddress     string `json:"street_address"`
	State             string `json:"state"`
	Zip               string `json:"zip"`
	GeneralPhone      string `json:"general_phone"`
	DirectorOfService string `json:"director_of_service"`
	LogoURL           any    `json:"logo_url"`
}

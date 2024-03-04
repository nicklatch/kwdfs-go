package domain

type Location struct {
	Id                   string `json:"id"`
	DealerName           string `json:"dealer"`
	BranchName           string `json:"branch_name"`
	State                string `json:"state"`
	City                 string `json:"city_county"`
	FleetSupportRep      string `json:"fleet_support_rep"`
	FleetSupportRepEmail string `json:"fleet_support_rep_email"`
	GeneralPhone         string `json:"general_phone"`
}

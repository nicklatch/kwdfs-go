package domain

type Customer struct {
	Id                   string `json:"id"`
	SponsoringDealer     string `json:"dealer"`
	Name                 string `json:"name"`
	State                string `json:"state"`
	PfleetAcctId         string `json:"pfleet_acct_id"`
	Status               string `json:"status"`
	TruckQty             int    `json:"truck_qty"`
	FleetSupportRep      string `json:"fleet_support_rep"`
	FieldServiceRep      string `json:"field_service_rep"`
	FieldServiceRepEmail string `json:"field_service_rep_email"`
}

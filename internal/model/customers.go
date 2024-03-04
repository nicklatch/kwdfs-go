package model

import "database/sql"

type Customer struct {
	Id                   string         `json:"id"`
	SponsoringDealer     string         `json:"dealer"`
	Name                 string         `json:"name"`
	State                string         `json:"state"`
	PfleetAcctId         sql.NullString `json:"pfleet_acct_id"`
	Status               string         `json:"status"`
	TruckQty             int            `json:"truck_qty"`
	FleetSupportRep      string         `json:"fleet_support_rep"`
	FieldServiceRep      sql.NullString `json:"field_service_rep"`
	FieldServiceRepEmail sql.NullString `json:"field_service_rep_email"`
}

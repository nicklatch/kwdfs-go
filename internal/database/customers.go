package database

import "log"

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

// GetAllCustomers is a Service method that returns
// either a slice of all customers, or an error
func (s Service) GetAllCustomers() ([]Customer, error) {
	// see ./dealers for comments to explain what takes place below
	rows, err := s.db.Query(`
select
  d.dealer, c.name, c.state, c.pfleet_acct_id, c.status, c.truck_qty,
  c.fleet_support_rep, c.field_service_rep, c.field_service_rep_email
from
  customers c
  inner join dealers d on c.dealer = d.id
  `)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.SponsoringDealer, &c.Name, &c.State, &c.PfleetAcctId,
			&c.Status, &c.TruckQty, &c.FleetSupportRep, &c.FieldServiceRep, &c.FieldServiceRepEmail); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

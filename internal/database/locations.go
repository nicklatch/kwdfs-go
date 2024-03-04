package database

import (
	"log"
	"nicklatch/kwdfs-go/internal/domain"
)

// GetAllLocations is a Service method that returns
// either a slice of all locations, or an error
func (s Service) GetAllLocations() ([]domain.Location, error) {
	// see ./dealers for comments to explain what takes place below
	rows, err := s.db.Query(`
select
  d.dealer,l.branch_name name,l.city_county,l.state,
  l.fleet_support_rep,l.fleet_support_rep_email,l.general_phone
from
  locations l
  inner join dealers d on l.dealer = d.id
  `)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var locations []domain.Location

	for rows.Next() {
		var l domain.Location
		if err := rows.Scan(&l.DealerName, &l.BranchName, &l.State, &l.City,
			&l.FleetSupportRep, &l.FleetSupportRepEmail, &l.GeneralPhone); err != nil {
			return nil, err
		}
		locations = append(locations, l)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

// TODO: Pagination
// TODO: pass args into queries instead of hard code (lines 23-28)

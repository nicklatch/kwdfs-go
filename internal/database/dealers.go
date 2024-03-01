package database

import (
	"log"
)

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

// GetAllDealerGroups is a Service method that returns either a
// slice of all dealers, or an error
func (s Service) GetAllDealerGroups() ([]Dealer, error) {
	// make db query and check for error
	rows, err := s.db.Query("SELECT * FROM dealers")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// The row reader will stay open by default, so we
	// defer its closing until the end of this scope
	defer rows.Close()

	var dealers []Dealer

	// iterates through our row results, unmarshalls them into the dealers slice above
	for rows.Next() {
		var d Dealer
		if err := rows.Scan(&d.Id, &d.DealerName, &d.StreetAddress, &d.State, &d.Zip,
			&d.GeneralPhone, &d.DirectorOfService, &d.LogoURL); err != nil {
			return nil, err
		}
		dealers = append(dealers, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dealers, nil
}

package db

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

// these will just read from json file until I can run local docker db's

func (s *service) GetAllDealerGroups() []Dealer {
	//file, err := os.ReadFile("internal/db/dealerDB.json")
	//if err != nil {
	//	log.Println(err)
	//}

	q := `select ?, ?, ?, ? from dealers`
	rows, err := s.db.Query(q, "dealer", "street_address", "state", "zip", "general_phone")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var dealers []Dealer

	//left off here need to get rows into proper type
	//err = json.Unmarshal(, &dealers)
	//if err != nil {
	//	log.Println(err)
	//}

	return dealers
}

package db

import (
	"encoding/json"
	"log"
	"os"
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

func GetAllDealerGroups() []Dealer {
	file, err := os.ReadFile("internal/db/dealerDB.json")
	if err != nil {
		log.Println(err)
	}
	var dealers []Dealer

	err = json.Unmarshal(file, &dealers)
	if err != nil {
		log.Println(err)
	}

	return dealers
}

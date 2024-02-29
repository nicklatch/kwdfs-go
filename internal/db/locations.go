package db

import (
	"encoding/json"
	"log"
	"os"
)

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

// TODO: Needs pagination

func GetAllLocations() []Location {
	file, err := os.ReadFile("internal/db/locationsDB.json")
	if err != nil {
		log.Println(err)
	}
	var locations []Location

	err = json.Unmarshal(file, &locations)
	if err != nil {
		log.Println(err)
	}

	return locations
}

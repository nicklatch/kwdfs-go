// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type FleetStatus string

const (
	FleetStatusRegional FleetStatus = "Regional"
	FleetStatusNational FleetStatus = "National"
)

func (e *FleetStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FleetStatus(s)
	case string:
		*e = FleetStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for FleetStatus: %T", src)
	}
	return nil
}

type NullFleetStatus struct {
	FleetStatus FleetStatus `json:"fleet_status"`
	Valid       bool        `json:"valid"` // Valid is true if FleetStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFleetStatus) Scan(value interface{}) error {
	if value == nil {
		ns.FleetStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FleetStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFleetStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FleetStatus), nil
}

type UserPosition string

const (
	UserPositionFleetRep        UserPosition = "fleet_rep"
	UserPositionServiceManager  UserPosition = "service_manager"
	UserPositionServiceDirector UserPosition = "service_director"
	UserPositionDealerPrincipal UserPosition = "dealer_principal"
	UserPositionStoreManager    UserPosition = "store_manager"
	UserPositionOther           UserPosition = "other"
)

func (e *UserPosition) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserPosition(s)
	case string:
		*e = UserPosition(s)
	default:
		return fmt.Errorf("unsupported scan type for UserPosition: %T", src)
	}
	return nil
}

type NullUserPosition struct {
	UserPosition UserPosition `json:"user_position"`
	Valid        bool         `json:"valid"` // Valid is true if UserPosition is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserPosition) Scan(value interface{}) error {
	if value == nil {
		ns.UserPosition, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserPosition.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserPosition) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserPosition), nil
}

type Customer struct {
	ID                   pgtype.UUID     `json:"id"`
	Dealer               pgtype.UUID     `json:"dealer"`
	Name                 string          `json:"name"`
	State                string          `json:"state"`
	PfleetAcctID         pgtype.Text     `json:"pfleet_acct_id"`
	Status               NullFleetStatus `json:"status"`
	TruckQty             int16           `json:"truck_qty"`
	FleetSupportRep      string          `json:"fleet_support_rep"`
	FieldServiceRep      pgtype.Text     `json:"field_service_rep"`
	FieldServiceRepEmail pgtype.Text     `json:"field_service_rep_email"`
}

type Dealer struct {
	ID     pgtype.UUID `json:"id"`
	Dealer string      `json:"dealer"`
	// Full address
	StreetAddress string `json:"street_address"`
	// Two letter state abbr
	State string      `json:"state"`
	Zip   pgtype.Text `json:"zip"`
	// XXX-XXX-XXXX format
	GeneralPhone string `json:"general_phone"`
	// First and last name
	DirectorOfService pgtype.Text `json:"director_of_service"`
	// url for s3
	LogoUrl   pgtype.Text        `json:"logo_url"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Location struct {
	ID         string      `json:"id"`
	Dealer     pgtype.UUID `json:"dealer"`
	BranchName string      `json:"branch_name"`
	// Two letter state abbr
	State      string `json:"state"`
	CityCounty string `json:"city_county"`
	// The name of the fleet support rep
	FleetSupportRep string `json:"fleet_support_rep"`
	// Email of the fleet support representitive
	FleetSupportRepEmail pgtype.Text `json:"fleet_support_rep_email"`
	// main phone number for a location
	GeneralPhone pgtype.Text `json:"general_phone"`
}

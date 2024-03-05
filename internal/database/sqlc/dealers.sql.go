// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: dealers.sql

package database

import (
	"context"
)

const getAllDealers = `-- name: GetAllDealers :many
SELECT id, dealer, street_address, state, zip, general_phone, director_of_service, logo_url FROM dealers
`

func (q *Queries) GetAllDealers(ctx context.Context) ([]Dealer, error) {
	rows, err := q.db.Query(ctx, getAllDealers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dealer
	for rows.Next() {
		var i Dealer
		if err := rows.Scan(
			&i.ID,
			&i.Dealer,
			&i.StreetAddress,
			&i.State,
			&i.Zip,
			&i.GeneralPhone,
			&i.DirectorOfService,
			&i.LogoUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDealerByName = `-- name: GetDealerByName :one
SELECT id, dealer, street_address, state, zip, general_phone, director_of_service, logo_url FROM dealers
WHERE dealer = $1
`

func (q *Queries) GetDealerByName(ctx context.Context, dealer string) (Dealer, error) {
	row := q.db.QueryRow(ctx, getDealerByName, dealer)
	var i Dealer
	err := row.Scan(
		&i.ID,
		&i.Dealer,
		&i.StreetAddress,
		&i.State,
		&i.Zip,
		&i.GeneralPhone,
		&i.DirectorOfService,
		&i.LogoUrl,
	)
	return i, err
}

const getDealersByState = `-- name: GetDealersByState :many
SELECT id, dealer, street_address, state, zip, general_phone, director_of_service, logo_url FROM dealers
WHERE state = $1
`

func (q *Queries) GetDealersByState(ctx context.Context, state string) ([]Dealer, error) {
	rows, err := q.db.Query(ctx, getDealersByState, state)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dealer
	for rows.Next() {
		var i Dealer
		if err := rows.Scan(
			&i.ID,
			&i.Dealer,
			&i.StreetAddress,
			&i.State,
			&i.Zip,
			&i.GeneralPhone,
			&i.DirectorOfService,
			&i.LogoUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
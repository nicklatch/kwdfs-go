-- name: GetAllDealers :many
SELECT *
FROM dealers;

-- name: GetDealerByName :one
SELECT *
FROM dealers
WHERE dealer = $1;

-- name: GetDealersByState :many
SELECT *
FROM dealers
WHERE state = $1;

-- name: CreateDealer :one
INSERT INTO dealers (dealer, street_address, state, zip, general_phone, director_of_service, logo_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: GetAllDealers :many
SELECT * FROM dealers;

-- name: GetDealerByName :one
SELECT * FROM dealers
WHERE dealer = $1;

-- name: GetDealersByState :many
SELECT * FROM dealers
WHERE state = $1;


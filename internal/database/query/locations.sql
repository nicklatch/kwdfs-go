-- name: GetAllLocations :many
SELECT d.dealer,
       l.branch_name,
       l.city_county,
       l.state,
       l.fleet_support_rep,
       l.fleet_support_rep_email,
       l.general_phone
FROM locations l
         INNER JOIN dealers d ON l.dealer = d.id
ORDER BY l.city_county;

-- name: GetLocationByBranchName :one
SELECT d.dealer,
       l.branch_name,
       l.city_county,
       l.state,
       l.fleet_support_rep,
       l.fleet_support_rep_email,
       l.general_phone
FROM locations l
         INNER JOIN dealers d ON l.dealer = d.id
WHERE l.branch_name = $1;
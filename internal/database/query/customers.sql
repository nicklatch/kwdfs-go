-- name: GetAllCustomers :many
SELECT d.dealer,
       c.name,
       c.state,
       c.pfleet_acct_id,
       c.status,
       c.truck_qty,
       c.fleet_support_rep,
       c.field_service_rep,
       c.field_service_rep_email
FROM customers c
         INNER JOIN dealers d ON c.dealer = d.id;

-- name: GetOneCustomer :one
SELECT d.dealer,
       c.name,
       c.state,
       c.pfleet_acct_id,
       c.status,
       c.truck_qty,
       c.fleet_support_rep,
       c.field_service_rep,
       c.field_service_rep_email
FROM customers c
         INNER JOIN dealers d ON c.dealer = d.id
WHERE c.name = $1;

-- name: CreateCustomer :one
INSERT INTO customers (dealer, name, state, pfleet_acct_id, truck_qty, fleet_support_rep, field_service_rep,
                       field_service_rep_email)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET name                    = $2,
    state                   = $3,
    pfleet_acct_id          = $4,
    status                  = $5,
    truck_qty               = $6,
    fleet_support_rep       = $7,
    field_service_rep       = $8,
    field_service_rep_email = $9
WHERE id = $1
RETURNING *;



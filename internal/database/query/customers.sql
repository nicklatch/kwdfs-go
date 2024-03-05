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
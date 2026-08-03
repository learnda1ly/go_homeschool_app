-- name: CreateFamily :one
INSERT INTO families (
  family_name
) VALUES (
  $1
) RETURNING *;

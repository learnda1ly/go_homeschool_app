-- name: CreateUser :one
INSERT INTO learners (
  family_id,
  first_name,
  last_name,
  grade
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

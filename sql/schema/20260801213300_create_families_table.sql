-- +goose Up
CREATE TABLE families(
  id SERIAL PRIMARY KEY,
  family_name VARCHAR(250) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS families;

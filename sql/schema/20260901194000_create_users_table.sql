-- +goose Up
CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(25) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL UNIQUE,
  family_id INTEGER NOT NULL REFERENCES families(id) ON DELETE CASCADE,
  role_id varchar(50) NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  encrypted_password VARCHAR(250) NOT NULL,
);

-- +goose Down
DROP TABLE IF EXISTS users;

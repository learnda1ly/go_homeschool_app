
-- +goose Up
CREATE TABLE learners(
  id SERIAL PRIMARY KEY,
  family_id INTEGER NOT NULL REFERENCES families(id),
  first_name VARCHAR(25) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  grade VARCHAR(10),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  image_url VARCHAR(250),
  learner_type VARCHAR(250)
);

-- +goose Down
DROP TABLE IF EXISTS learners;

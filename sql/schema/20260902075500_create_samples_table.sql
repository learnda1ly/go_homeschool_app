-- +goose Up
CREATE TABLE samples(
  id SERIAL PRIMARY KEY,
  learner_id INTEGER NOT NULL REFERENCES learners(id) ON DELETE CASCADE,
  subject_id INTEGER NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
  title VARCHAR(250) NOT NULL,
  description VARCHAR(2500),
  notes varchar(2500),
  url VARCHAR(250),
  letter_grade VARCHAR(5),
  percentage_grade int,
  completion_time TIMESTAMP,
  scheduled_time TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS samples;

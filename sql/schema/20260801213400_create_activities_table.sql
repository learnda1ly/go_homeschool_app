
-- +goose Up
CREATE TABLE activities(
  id INTEGER PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  description VARCHAR(2500),
  purpose VARCHAR(2500),
  scheduled_start_time TIMESTAMP,
  start_time TIMESTAMP,
  end_time TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS activities;

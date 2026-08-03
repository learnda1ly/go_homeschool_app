
-- +goose Up
CREATE TABLE activities(
  id INTEGER PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  description VARCHAR(2500),
  purpose VARCHAR(2500),
  start_date_time TIMESTAMP,
  end_date_time TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS activities;

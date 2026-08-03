
-- +goose Up
CREATE TABLE activities_learners(
  id INTEGER PRIMARY KEY,
  learner_id INTEGER NOT NULL REFERENCES learners(id) ON DELETE CASCADE,
  activity_id INTEGER NOT NULL REFERENCES activities(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS activities_learners;

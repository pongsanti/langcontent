-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE mcontents (
  id serial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  label text NOT NULL DEFAULT '',
  content_type text NOT NULL DEFAULT 'news',
  start_at timestamp with time zone,
  end_at timestamp with time zone,
  status text,
  xtime1 timestamp with time zone,
  creator_id integer
);

CREATE INDEX idx_mcontents_content_type ON mcontents(content_type);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE mcontents;

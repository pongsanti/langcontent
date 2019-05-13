-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE mcontent_texts (
  id serial PRIMARY KEY,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  lang text NOT NULL DEFAULT 'en',
  title text NOT NULL,
  subtitle text,
  detail text,
  xtext1 text,
  creator_id integer,
  mcontent_id integer NOT NULL REFERENCES mcontents(id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE INDEX idx_mcontent_texts_lang ON mcontent_texts(lang);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE mcontent_texts;

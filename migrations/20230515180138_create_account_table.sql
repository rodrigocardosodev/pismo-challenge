-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS accounts (
  "id" SERIAL PRIMARY KEY, 
  "document_number" VARCHAR(11) NOT NULL UNIQUE, 
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE INDEX document_number_idx ON accounts (document_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd

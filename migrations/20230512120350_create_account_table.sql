-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS accounts ("id" SERIAL NULL PRIMARY KEY, "document_number" VARCHAR(11) NOT NULL UNIQUE, "created_at" TIMESTAMP NOT NULL);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS accounts;

-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS accounts ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "document_number" varchar(11) NOT NULL UNIQUE, "created_at" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS accounts;

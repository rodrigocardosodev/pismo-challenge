-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS transactions (
  "id" SERIAL PRIMARY KEY, 
  "operation_id" SMALLINT NOT NULL, 
  "amount" FLOAT NOT NULL, 
  "event_date" TIMESTAMP NOT NULL DEFAULT NOW(), 
  "account_id" integer NOT NULL, 
  FOREIGN KEY (account_id) REFERENCES accounts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd

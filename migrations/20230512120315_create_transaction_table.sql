-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS transactions ("id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, "operation_id" integer NOT NULL, "amount" integer NOT NULL, "event_date" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, "account_id" integer NOT NULL, FOREIGN KEY (account_id) REFERENCES accounts(id));

CREATE TRIGGER verify_account_before_insert BEFORE INSERT ON transactions FOR EACH ROW BEGIN SELECT CASE WHEN ((SELECT id FROM accounts WHERE id = NEW.account_id) IS NULL) THEN RAISE (ABORT, 'Account not found') END; END;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS transactions;
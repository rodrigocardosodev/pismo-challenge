-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS transactions ("id" SERIAL NOT NULL PRIMARY KEY, "operation_id" SMALLINT NOT NULL, "amount" BIGINT NOT NULL, "event_date" TIMESTAMP NOT NULL, "account_id" integer NOT NULL, FOREIGN KEY (account_id) REFERENCES accounts(id));

CREATE TRIGGER verify_account_before_insert BEFORE INSERT ON transactions FOR EACH ROW BEGIN SELECT CASE WHEN ((SELECT id FROM accounts WHERE id = NEW.account_id) IS NULL) THEN RAISE (ABORT, 'Account not found') END; END;

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS transactions;
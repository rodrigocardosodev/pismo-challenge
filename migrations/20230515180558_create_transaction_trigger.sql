-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL trigger query';

CREATE OR REPLACE FUNCTION verify_account_exists() RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM accounts WHERE id = NEW.account_id) THEN
        RAISE EXCEPTION 'account not found';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_account_exists
BEFORE INSERT ON transactions
FOR EACH ROW
EXECUTE PROCEDURE verify_account_exists();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL trigger query';

DROP TRIGGER IF EXISTS check_account_exists ON transactions;
DROP FUNCTION IF EXISTS verify_account_exists();
-- +goose StatementEnd

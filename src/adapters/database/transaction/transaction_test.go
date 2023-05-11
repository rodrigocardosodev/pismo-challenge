package transaction_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import SQLite3 driver

	database "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/transaction"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/stretchr/testify/require"
)

var transactionDb *sql.DB

func initDB() {
	var err error
	transactionDb, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal("Failed to initialize database:", err.Error())
	}
}

func setUp() {
	initDB()
	createAccountsTable(transactionDb)
	createTransactionsTable(transactionDb)
}

func tearDown(db *sql.DB) {
	query := "DROP TABLE transactions; DROP TABLE accounts;"
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec()
}

func createAccountsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS accounts (
							"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"document_number" varchar(11) NOT NULL,
							"amount" integer NOT NULL
						);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createTransactionsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS transactions (
							"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"operation_id" integer NOT NULL,
							"amount" integer NOT NULL,
							"event_date" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
							"account_id" integer NOT NULL,
							FOREIGN KEY (account_id) REFERENCES accounts(id)
    				);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createAccount(db *sql.DB) {
	query := `INSERT INTO accounts (document_number, amount) VALUES ("55724203014", 1000);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestTransactionRepository_Create(t *testing.T) {
	setUp()
	createAccount(transactionDb)
	defer tearDown(transactionDb)
	ctx := context.Background()

	transation := models.NewTransaction(1, models.PAGAMENTO, 1000)
	transactionDb := database.NewTransactionRepository(transactionDb)

	result, err := transactionDb.Create(ctx, transation)
	require.Nil(t, err)
	require.Equal(t, transation, result)
}
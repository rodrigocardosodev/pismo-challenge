package account_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import SQLite3 driver
	database "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/account"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/stretchr/testify/require"
)

var accountDb *sql.DB

func initDB() {
	var err error
	accountDb, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal("Failed to initialize database:", err.Error())
	}
}

func setUp() {
	initDB()
	createTable(accountDb)
}

func tearDown(db *sql.DB) {
	query := "DROP TABLE accounts;"
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	stmt.Exec()
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS accounts (
							"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
							"document_number" varchar(11) NOT NULL UNIQUE
    				);`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createAccount(db *sql.DB) {
	query := `INSERT INTO accounts (document_number) VALUES ("55724203014");`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestAccountRepository_GetById(t *testing.T) {
	setUp()
	createAccount(accountDb)
	defer tearDown(accountDb)
	ctx := context.Background()

	accountDb := database.NewAccountRepository(accountDb)
	account, err := accountDb.GetByID(ctx, 1)
	require.Nil(t, err)
	require.Equal(t, int64(1), account.GetID())
	require.Equal(t, "55724203014", account.GetDocumentNumber())

	err = account.IsValid()
	require.Nil(t, err)
}

func TestAccountRepository_Create(t *testing.T) {
	setUp()
	defer tearDown(accountDb)
	ctx := context.Background()

	account := models.NewAccount("55724203014")
	accountDb := database.NewAccountRepository(accountDb)
	account, err := accountDb.Create(ctx, account)
	require.Nil(t, err)
	require.Equal(t, int64(1), account.GetID())
	require.Equal(t, "55724203014", account.GetDocumentNumber())

	err = account.IsValid()
	require.Nil(t, err)
	accountAlreadyExists := models.NewAccount("55724203014")
	account, err = accountDb.Create(ctx, accountAlreadyExists)
	require.NotNil(t, err)
	require.Nil(t, account)
}

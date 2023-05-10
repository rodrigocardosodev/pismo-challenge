package database_test

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import SQLite3 driver
	"github.com/rodrigocardosodev/pismo-challenge/src/adapters/database"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/models"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func initDB() {
	var err error
	Db, err = sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatal("Failed to initialize database:", err.Error())
	}
}

func setUp() {
	initDB()
	createTable(Db)
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
							"document_number" varchar(11) NOT NULL,
							"amount" integer NOT NULL
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

func TestAccountRepository_GetById(t *testing.T) {
	setUp()
	createAccount(Db)
	defer tearDown(Db)

	accountDb := database.NewAccountRepository(Db)
	account, err := accountDb.GetByID(1)
	require.Nil(t, err)
	require.Equal(t, int64(1), account.GetID())
	require.Equal(t, "55724203014", account.GetDocumentNumber())
	require.Equal(t, int64(1000), account.GetAmount())
}

func TestAccountRepository_Create(t *testing.T) {
	setUp()
	defer tearDown(Db)

	account := models.NewAccount("55724203014")
	accountDb := database.NewAccountRepository(Db)
	account, err := accountDb.Create(account)
	require.Nil(t, err)
	require.Equal(t, int64(1), account.GetID())
	require.Equal(t, "55724203014", account.GetDocumentNumber())
	require.Equal(t, int64(0), account.GetAmount())
}

func TestAccountRepository_UpdateAmount(t *testing.T) {
	setUp()
	createAccount(Db)
	defer tearDown(Db)

	accountDb := database.NewAccountRepository(Db)
	account, err := accountDb.UpdateAmount(1, 2000)
	require.Nil(t, err)
	require.Equal(t, int64(1), account.GetID())
	require.Equal(t, "55724203014", account.GetDocumentNumber())
	require.Equal(t, int64(2000), account.GetAmount())
}
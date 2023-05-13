package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	userDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/account"
	transactionDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/transaction"

	userHTTP "github.com/rodrigocardosodev/pismo-challenge/src/adapters/http/account"
	transactionHTTP "github.com/rodrigocardosodev/pismo-challenge/src/adapters/http/transaction"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"

	// import postgres driver
	_ "github.com/lib/pq"
)

func main() {
	var (
		DB_HOST     = os.Getenv("PG_HOST")
		DB_PORT     = os.Getenv("PG_PORT")
		DB_USER     = os.Getenv("PG_USER")
		DB_PASSWORD = os.Getenv("PG_PASSWORD")
		DB_DATABASE = os.Getenv("PG_DATABASE")
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)
	db, err := sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountRepository := userDatabase.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	accountRoutes := userHTTP.NewHTTPAccountAdapter(accountService)

	transactionRepository := transactionDatabase.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionRoutes := transactionHTTP.NewHTTPTransactionAdapter(transactionService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())

	router.POST("/accounts", accountRoutes.CreateAccount)
	router.GET("/accounts/:account_id", accountRoutes.GetAccountById)
	router.POST("/transactions", transactionRoutes.CreateTransaction)

	router.Run(":8080")
}

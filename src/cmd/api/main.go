package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	userDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/account"
	transactionDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/database/transaction"

	userHTTP "github.com/rodrigocardosodev/pismo-challenge/src/adapters/http/account"
	transactionHTTP "github.com/rodrigocardosodev/pismo-challenge/src/adapters/http/transaction"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"

	// import sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./pismo.db")
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

	router.POST("/accounts", accountRoutes.CreateAccount)
	router.GET("/accounts/:account_id", accountRoutes.GetAccountById)
	router.POST("/transactions", transactionRoutes.CreateTransaction)

	router.Run(":8080")
}

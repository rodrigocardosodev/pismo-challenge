package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	userDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/secondary/database/postgres/account"
	transactionDatabase "github.com/rodrigocardosodev/pismo-challenge/src/adapters/secondary/database/postgres/transaction"

	userHTTPAdapter "github.com/rodrigocardosodev/pismo-challenge/src/adapters/primary/http/account"
	transactionHTTPAdapter "github.com/rodrigocardosodev/pismo-challenge/src/adapters/primary/http/transaction"
	"github.com/rodrigocardosodev/pismo-challenge/src/application/services"

	_ "github.com/lib/pq"
	_ "github.com/rodrigocardosodev/pismo-challenge/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Pismo Challenge API
//	@version		1.0
//	@description	This is a challenge for Pismo's backend developer position.

//	@contact.name	Rodrigo Cardoso
//	@contact.email	rodrigo.cardoso92@outlook.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	var (
		DB_HOST     = os.Getenv("POSTGRES_HOST")
		DB_PORT     = os.Getenv("POSTGRES_PORT")
		DB_USER     = os.Getenv("POSTGRES_USER")
		DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
		DB_DATABASE = os.Getenv("POSTGRES_DB")
	)

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)
	db, err := sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	transactionRepository := transactionDatabase.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionRoutes := transactionHTTPAdapter.NewHTTPTransactionAdapter(transactionService)

	accountRepository := userDatabase.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository, transactionRepository)
	accountRoutes := userHTTPAdapter.NewHTTPAccountAdapter(accountService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/accounts", accountRoutes.CreateAccount)
	router.GET("/accounts/:account_id", accountRoutes.GetAccountById)
	router.GET("/accounts/:account_id/balance", accountRoutes.GetAccountBalance)

	router.POST("/transactions", transactionRoutes.CreateTransaction)

	router.Run(":8080")
}

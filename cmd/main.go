package main

import (
	"database/sql"
	"fmt"

	"github.com/flvSantos15/dtmoney-transactions/internal/database"
	"github.com/flvSantos15/dtmoney-transactions/internal/usecase/create_transaction"
	"github.com/flvSantos15/dtmoney-transactions/internal/web"
	"github.com/flvSantos15/dtmoney-transactions/internal/web/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dtmoney")
	if err != nil {
		fmt.Println("error aqui =>", err)
		panic(err.Error())
	}
	defer db.Close()

	transactionDB := database.NewTransactionDB(db)

	createTransactionUseCase := create_transaction.NewTransactionUseCase(transactionDB)

	webserver := webserver.NewWebServer(":8080")

	transactionHandler := web.NewTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webserver.Start()
}
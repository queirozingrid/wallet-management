package main

import (
	"fmt"
	"wallet-management/database"
	"wallet-management/types"

	_ "github.com/lib/pq"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	transaction := types.Transaction{
		Amount:      35.50,
		Description: "'Something nice here'",
		Type:        types.INFLOW,
	}

	err := database.CreateTransaction(db, transaction)
	if err != nil {
		panic(err)
	}

	transactions, err := database.GetAllTransactions(db)
	if err != nil {
		panic(err)
	}

	for _, tr := range transactions {
		fmt.Printf("---Transaction: %+v - %+v - %+v - %+v ---\n", tr.Id, tr.Amount, tr.Description, tr.Type)
	}

}

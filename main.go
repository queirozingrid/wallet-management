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

	// CREATE A NEW TRANSACTION
	fmt.Println("Creating transaction...")
	err := database.CreateTransaction(db, transaction)
	if err != nil {
		panic(err)
	}

	// GET ALL TRANSACTIONS IN THE DB
	fmt.Println("All transactions in the Database...")
	transactions, err := database.GetAllTransactions(db)
	if err != nil {
		panic(err)
	}

	for _, tr := range transactions {
		fmt.Printf("---Transaction: %+v - %+v - %+v - %+v ---\n", tr.Id, tr.Amount, tr.Description, tr.Type)
	}

	updateTransaction := types.Transaction{
		Id:          1,
		Amount:      56.80,
		Description: "'Something bad here'",
		Type:        types.OUTFLOW,
	}

	// UPDATE THIS ONE TRANSACTION
	fmt.Println("\nUpdating transaction...")
	err = database.UpdateTransaction(db, updateTransaction)
	if err != nil {
		panic(err)
	}

	// GET ALL TRANSACTIONS IN THE DB
	fmt.Println("All transactions in the Database...")
	transactions, err = database.GetAllTransactions(db)
	if err != nil {
		panic(err)
	}

	for _, tr := range transactions {
		fmt.Printf("---Transaction: %+v - %+v - %+v - %+v ---\n", tr.Id, tr.Amount, tr.Description, tr.Type)
	}

	// DELETE THIS ONE TRANSACTION
	fmt.Println("\nDeleting transaction...")
	err = database.DeleteTransactionById(db, 1)
	if err != nil {
		panic(err)
	}

	// GET ALL TRANSACTIONS IN THE DB
	fmt.Println("All transactions in the Database...")
	transactions, err = database.GetAllTransactions(db)
	if err != nil {
		panic(err)
	}

	for _, tr := range transactions {
		fmt.Printf("---Transaction: %+v - %+v - %+v - %+v ---\n", tr.Id, tr.Amount, tr.Description, tr.Type)
	}

}

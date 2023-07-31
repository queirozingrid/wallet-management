package database

import (
	"database/sql"
	"fmt"
	"wallet-management/types"
)

func CreateTransaction(db *sql.DB, transaction types.Transaction) (err error) {

	query := fmt.Sprintf("INSERT INTO transactions"+
		"(amount, description, transaction_type)"+
		"VALUES (%+v, %+v, %+v)", transaction.Amount, transaction.Description, transaction.Type)
	_, err = db.Query(query)
	if err != nil {
		return (err)
	}
	return nil
}

func GetAllTransactions(db *sql.DB) (transactions []types.Transaction, err error) {
	query := fmt.Sprintf("SELECT * FROM transactions")
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var transaction types.Transaction
		err = rows.Scan(&transaction.Id, &transaction.Amount, &transaction.Description, &transaction.Type)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func UpdateTransaction(db *sql.DB, transaction types.Transaction) (err error) {
	query := fmt.Sprintf("UPDATE transactions "+
		"SET amount=%+v, description=%+v, transaction_type=%+v "+
		"WHERE id=%+v", transaction.Amount, transaction.Description, transaction.Type, transaction.Id)

	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTransactionById(db *sql.DB, id int) (err error) {
	query := fmt.Sprintf("DELETE FROM transactions WHERE id=%+v", id)
	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

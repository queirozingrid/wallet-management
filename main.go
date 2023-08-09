package main

import (
	"wallet-management/database"

	_ "github.com/lib/pq"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()
}

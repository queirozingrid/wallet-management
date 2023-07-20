package database

import "database/sql"

func ConnectDB() *sql.DB {
	// Database Connection
	conn := "user=user_name dbname=db_name password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}

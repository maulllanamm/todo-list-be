package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func NewConnection() *sql.DB {
	connStr := "user=maulllanam password=qwe123 host=localhost port=5432 dbname=todo-list-golang sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil{
		panic(err)
	}
	
	return db
}
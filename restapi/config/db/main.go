package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect : creates a connection to our database
func Connect() {
	db, err := sql.Open("mysql", "root:Amir2222@tcp(127.0.0.1:3306)/books")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to mysql successfully")

	DB = db
}

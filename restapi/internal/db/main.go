package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Amir2222@tcp(127.0.0.1:3306)/books")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to mysql successfully")

	defer db.Close()
}

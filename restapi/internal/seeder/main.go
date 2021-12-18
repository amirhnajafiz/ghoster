package seeder

import (
	"context"
	"restapi/restapi/config/db"
	"time"
)

func Seed() {
	connection := db.DB

	query := `CREATE TABLE IF NOT EXISTS book(name varchar(20) primary key)`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := connection.ExecContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}
}

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restapi/restapi/internal/database/migrations"
	"restapi/restapi/internal/database/seeder"
)

var DB *gorm.DB

// Connect : creates a connection to our database
func Connect() {
	dsn := "root:Amir2222@tcp(127.0.0.1:3306)/books"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Migrating the database
	migrations.Migrate(db)
	seeder.Seed(db)

	DB = db
}

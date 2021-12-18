package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"restapi/restapi/internal/database/migrations"
)

var DB *gorm.DB

// Connect : creates a connection to our database
func Connect() {
	dsn := "root:Amir2222@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	// Migrating the database
	migrations.Migrate(db)

	DB = db
}

package migrations

import (
	"gorm.io/gorm"
	"restapi/restapi/internal/database/models/book"
)

// Migrate : does the migrations of database
func Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(&book.Book{})
}

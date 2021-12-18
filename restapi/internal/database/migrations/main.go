package migrations

import (
	"gorm.io/gorm"
	"restapi/restapi/internal/models/book"
)

// Migrate : does the migrations of database
func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&book.Book{})
}

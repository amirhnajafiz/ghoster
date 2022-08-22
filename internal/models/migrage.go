package models

import "gorm.io/gorm"

// Migrate : does the migrations of database
func Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(&Book{})
}

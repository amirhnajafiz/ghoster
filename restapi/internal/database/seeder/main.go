package seeder

import (
	"gorm.io/gorm"
	"math/rand"
	"restapi/restapi/internal/database/models/book"
	"strconv"
)

// Seed : puts the initialized data into the database
func Seed(db *gorm.DB) {
	db.Create(&book.Book{
		Isbn:   strconv.Itoa(rand.Int()),
		Title:  "First book",
		Author: "Mark Antony",
	})

	db.Create(&book.Book{
		Isbn:   strconv.Itoa(rand.Int()),
		Title:  "Second book",
		Author: "Mary Morphy",
	})
}

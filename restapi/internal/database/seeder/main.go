package seeder

import (
	"gorm.io/gorm"
	"restapi/restapi/internal/models/author"
	"restapi/restapi/internal/models/book"
)

// Seed : puts the initialized data into the database
func Seed(db *gorm.DB) {
	db.Create(&book.Book{
		ID:    "1",
		Isbn:  "129987",
		Title: "First book",
		Author: &author.Author{
			Firstname: "Mike",
			Lastname:  "Arte",
		},
	})

	db.Create(&book.Book{
		ID:    "2",
		Isbn:  "220199",
		Title: "Second book",
		Author: &author.Author{
			Firstname: "Mary",
			Lastname:  "Antoin",
		},
	})
}

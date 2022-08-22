package seeder

import (
	"gorm.io/gorm"
	"math/rand"
	"restapi/restapi/internal/database/models/author"
	"restapi/restapi/internal/database/models/book"
	"strconv"
)

// Seed : puts the initialized data into the database
func Seed(db *gorm.DB) {
	db.Create(&author.Author{
		Firstname: "Roman",
		Lastname:  "Petrovic",
	})

	db.Create(&author.Author{
		Firstname: "Sandy",
		Lastname:  "Jovic",
	})

	db.Create(&author.Author{
		Firstname: "Peter",
		Lastname:  "Lubernth",
	})

	db.Create(&author.Author{
		Firstname: "Andy",
		Lastname:  "Williams",
	})

	db.Create(&author.Author{
		Firstname: "Dandy",
		Lastname:  "Warhole",
	})

	db.Create(&author.Author{
		Firstname: "Kevin",
		Lastname:  "Ninto",
	})

	db.Create(&book.Book{
		Isbn:     strconv.Itoa(rand.Int()),
		Title:    "First book",
		AuthorID: 1,
	})

	db.Create(&book.Book{
		Isbn:     strconv.Itoa(rand.Int()),
		Title:    "Second book",
		AuthorID: 2,
	})
}

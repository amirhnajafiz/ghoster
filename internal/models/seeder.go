package models

import (
	"math/rand"
	"strconv"

	"gorm.io/gorm"
)

// Seed : puts the initialized data into the database
func Seed(db *gorm.DB) {
	db.Create(&Author{
		Firstname: "Roman",
		Lastname:  "Petrovic",
	})

	db.Create(&Author{
		Firstname: "Sandy",
		Lastname:  "Jovic",
	})

	db.Create(&Author{
		Firstname: "Peter",
		Lastname:  "Lubernth",
	})

	db.Create(&Author{
		Firstname: "Andy",
		Lastname:  "Williams",
	})

	db.Create(&Author{
		Firstname: "Dandy",
		Lastname:  "Warhole",
	})

	db.Create(&Author{
		Firstname: "Kevin",
		Lastname:  "Ninto",
	})

	db.Create(&Book{
		Isbn:     strconv.Itoa(rand.Int()),
		Title:    "First book",
		AuthorID: 1,
	})

	db.Create(&Book{
		Isbn:     strconv.Itoa(rand.Int()),
		Title:    "Second book",
		AuthorID: 2,
	})
}

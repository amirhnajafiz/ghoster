package book

import (
	"gorm.io/gorm"
)

// Book struct (Model)
type Book struct {
	gorm.Model
	Isbn   string
	Title  string
	Author string
}

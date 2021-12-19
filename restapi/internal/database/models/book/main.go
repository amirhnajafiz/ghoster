package book

import "restapi/restapi/internal/database/models/author"

// Book struct (Model)
type Book struct {
	ID       int           `gorm:"primary_key" json:"id"`
	Isbn     string        `json:"isbn"`
	Title    string        `json:"title"`
	AuthorID int           `json:"author_id"`
	Author   author.Author `gorm:"references:ID"`
}

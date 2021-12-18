package book

import A "restapi/restapi/internal/models/author"

// Book struct (Model)
type Book struct {
	ID     string    `json:"id"`
	Isbn   string    `json:"isbn"`
	Title  string    `json:"title"`
	Author *A.Author `json:"author"`
}

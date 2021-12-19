package book

// Book struct (Model)
type Book struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

package author

// Author struct
type Author struct {
	ID        int
	Firstname string `gorm:"primaryKey"`
	Lastname  string
}

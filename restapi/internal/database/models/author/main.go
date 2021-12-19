package author

// Author struct
type Author struct {
	ID        int `gorm:"primaryKey"`
	Firstname string
	Lastname  string
}

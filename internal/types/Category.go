package types

// Category represents the category entity in the database.
type Category struct {
	CategoryID  uint   `gorm:"primaryKey"` // Primary key
	Name        string `gorm:"not null"`   // Category name
	Description string `gorm:"type:text"`  // Description of the category
}

package types



// Manufacturer represents the manufacturer entity in the database.
type Manufacturer struct {
	ManufacturerID uint   `gorm:"primaryKey"` // Primary key
	Name           string `gorm:"not null"`   // Manufacturer name
	Description    string `gorm:"type:text"`  // Description of the manufacturer

	// Associations
	Products []Product `gorm:"foreignKey:ManufacturerID"` // One-to-many with Product
}

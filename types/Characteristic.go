package types



// Characteristic represents the characteristic entity in the database.
type Characteristic struct {
	CharacteristicID uint   `gorm:"primaryKey"` // Primary key
	Name             string `gorm:"not null"`   // Name of the characteristic
	CategoryID       uint   `gorm:"not null"`   // Foreign key to Category

	// Associations
	Category Category  `gorm:"foreignKey:CategoryID"`             // Belongs to Category
	Products []Product `gorm:"many2many:product_characteristics"` // Many-to-many with Product
}

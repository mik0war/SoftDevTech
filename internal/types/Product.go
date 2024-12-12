package types

// Product represents the product entity in the database.
type Product struct {
	ProductID      uint   `gorm:"primaryKey" json:"product_id"`       // Primary key
	Name           string `gorm:"not null" json:"name"`               // Product name
	Status         string `gorm:"not null" json:"status"`             // Status
	Description    string `gorm:"type:text" json:"description"`       // Product description
	ImageLink      string `gorm:"type:varchar(255)" json:"imageLink"` // URL to the product image
	ManufacturerID uint   `gorm:"not null" json:"manufacturerID"`     // Foreign key to Manufacturer

	// Associations
	Manufacturer   *Manufacturer    `gorm:"foreignKey:ManufacturerID;"`                                                               // Belongs to Manufacturer
	Category       []Category       `gorm:"many2many:product_category;joinForeignKey:ProductID;joinReferences:CategoryID"`            // Many-to-many with Category
	Characteristic []Characteristic `gorm:"many2many:ProductCharacteristic;joinForeignKey:ProductID;joinReferences:CharacteristicID"` // Many-to-many with Characteristic
}

package types

// Product represents the product entity in the database.
type Product struct {
	ProductID      uint   `gorm:"primaryKey;autoIncrement" json:"product_id"` // Primary key
	Name           string `gorm:"not null" json:"name"`                       // Product name
	Status         string `gorm:"not null" json:"status"`                     // Status
	Description    string `gorm:"type:text" json:"description"`               // Product description
	ImageLink      string `gorm:"type:varchar(255)" json:"imageLink"`         // URL to the product image
	ManufacturerID uint   `gorm:"not null" json:"manufacturer_id"`            // Foreign key to Manufacturer

	// Associations
	Manufacturer   *Manufacturer `gorm:"foreignKey:ManufacturerID;references:ManufacturerID"`
	Category       []ProductCategory
	Characteristic []ProductCharacteristic

	Cost []ProductCost
}

package types

// Characteristic represents the characteristic entity in the database.
type Characteristic struct {
	CharacteristicID int    `gorm:"primaryKey"` // Primary key
	Name             string `gorm:"not null"`   // Name of the characteristic
	Category         string `gorm:"not null"`   // Name of category
}

type ProductCharacteristic struct {
	ProductID        uint           `gorm:"primaryKey" json:"-"`
	CharacteristicID uint           `gorm:"primaryKey" json:"-"`
	Characteristic   Characteristic `gorm:"foreignKey:CharacteristicID;references:CharacteristicID"`
	Value            string         `gorm:"not null"`
}

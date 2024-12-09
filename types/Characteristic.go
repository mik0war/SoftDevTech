package types

// Characteristic represents the characteristic entity in the database.
type Characteristic struct {
	CharacteristicID uint   `gorm:"primaryKey"` // Primary key
	Name             string `gorm:"not null"`   // Name of the characteristic
	Category         string `gorm:"not null"`   // Name of category
	Value            string `gorm:"not null"`   // Value of the characteristic
}

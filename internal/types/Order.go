package types

import "time"

// Order — структура для хранения данных о заказе
type Order struct {
	OrderID   int       `gorm:"primaryKey" json:"order_id"` // Уникальный идентификатор заказа
	Status    string    `json:"status"`                     // Статус заказа
	TimeStamp time.Time `json:"time_stamp"`                 // Временная метка создания заказа

	Products []ProductOrder `gorm:"foreignKey:OrderID" json:"products"` // Список продуктов в заказе
}

// ProductOrder — структура для хранения данных о продукте в заказе
type ProductOrder struct {
	OrderID   int `gorm:"primaryKey" json:"order_id"`   // ID заказа (внешний ключ)
	ProductID int `gorm:"primaryKey" json:"product_id"` // ID продукта (внешний ключ)
	Count     int `json:"count"`                        // Количество продукта

	Product Product `gorm:"foreignKey:ProductID;references:ProductID" json:"product"`
}

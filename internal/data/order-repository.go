package data

import (
	"online-shop-API/internal/types"
)

func (repo *Repository) CreateOrder(order *types.Order) (int, error) {
	var orderID int
	query := `
        INSERT INTO "order" (status, time_stamp)
        VALUES ($1, $2)
        RETURNING order_id;
    `
	err := repo.db.Raw(query, order.Status, order.TimeStamp).Scan(&orderID).Error
	if err != nil {
		return 0, err
	}
	return orderID, nil
}

func (repo *Repository) AddProductToOrder(productOrder *types.ProductOrder) error {
	query := `
        INSERT INTO product_order (order_id, product_id, count)
        VALUES ($1, $2, $3);
    `
	err := repo.db.Exec(
		query, productOrder.OrderID, productOrder.ProductID, productOrder.Count,
	).Error

	return err
}

func (repo *Repository) GetOrderWithProducts(orderID uint) (*types.Order, error) {
	var order types.Order

	// Используем Preload для загрузки связанных данных из таблицы product_in_order
	err := repo.db.
		Preload("Products").
		Preload("Products.Product").
		Preload("Products.Product.Manufacturer").
		Preload("Products.Product.Category").
		Preload("Products.Product.Characteristic").
		Preload("Products.Product.Cost").
		First(&order, orderID).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

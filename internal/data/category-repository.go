package data

import (
	"errors"
	"online-shop-API/internal/types"
)

func (repo *Repository) SubscribeProductCategory(productId uint, categoryId uint) error {

	if err := repo.db.Create(types.ProductCategory{ProductID: productId, CategoryID: categoryId}); err != nil {
		return errors.New("missing productId or CategoryId")
	}

	return nil
}

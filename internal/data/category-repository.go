package data

import (
	"errors"
	"online-shop-API/internal/types"
)

func (repo *Repository) SubscribeProductCategory(productId uint, categoryId string) error {

	var category types.Category
	repo.db.First(&category, "name = ?", categoryId)
	if err := repo.db.Create(types.ProductCategory{ProductID: productId, CategoryID: category.CategoryID}).Error; err != nil {
		return errors.New("missing productId or CategoryId")
	}

	return nil
}

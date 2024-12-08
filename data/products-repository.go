package data

import (
	"errors"
	"gorm.io/gorm"
	"online-shop-API/types"
)

// ProductRepository provides access to the product store.
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetProducts retrieves a list of products with optional filters and pagination.
func (repo *ProductRepository) GetProducts(
	filters map[string]interface{},
	page int,
	pageSize int,
) ([]types.Product, int64, error) {

	var products []types.Product
	var totalCount int64

	// Initialize the query
	query := repo.db.Model(&types.Product{})

	// Apply filters
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// Count total products for pagination
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	_ = (page - 1) * pageSize
	if err := query.Preload("Manufacturer").Preload("Categories").Preload("Characteristics").
		Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

// AddProduct добавляет новый продукт в базу данных.
func (repo *ProductRepository) AddProduct(product *types.Product) (*types.Product, error) {
	// Проверка на обязательные поля
	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.ManufacturerID == 0 {
		return nil, errors.New("manufacturer ID is required")
	}

	// Сохраняем продукт в базе
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}

	// Возвращаем созданный продукт
	return product, nil
}

// UpdateProduct обновляет существующий продукт в базе данных.
func (repo *ProductRepository) UpdateProduct(
	productID uint,
	updatedData *types.Product,
) (*types.Product, error) {
	// Проверяем, существует ли продукт
	var product types.Product
	if err := repo.db.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Обновляем поля
	if updatedData.Name != "" {
		product.Name = updatedData.Name
	}
	if updatedData.Status != "" {
		product.Status = updatedData.Status
	}
	if updatedData.Description != "" {
		product.Description = updatedData.Description
	}
	if updatedData.ImageLink != "" {
		product.ImageLink = updatedData.ImageLink
	}
	if updatedData.ManufacturerID != 0 {
		product.ManufacturerID = updatedData.ManufacturerID
	}

	// Сохраняем изменения
	if err := repo.db.Save(&product).Error; err != nil {
		return nil, err
	}

	// Возвращаем обновленный продукт
	return &product, nil
}

// DeleteProduct удаляет продукт из базы данных по его ID.
func (repo *ProductRepository) DeleteProduct(productID uint) error {
	// Проверяем, существует ли продукт
	var product types.Product
	if err := repo.db.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	// Удаляем продукт
	if err := repo.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

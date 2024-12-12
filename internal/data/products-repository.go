package data

import (
	"errors"
	"gorm.io/gorm"
	"online-shop-API/internal/types"
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

	repo.db = repo.db.Debug()
	var products []types.Product
	var totalCount int64

	// Initialize the productsQuery
	productsQuery := repo.db.Model(&types.Product{})

	// Apply filters
	for key, value := range filters {
		productsQuery = productsQuery.Where(key+" = ?", value)
	}

	// Count total products for pagination
	if err := productsQuery.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	// Сначала получаем список продуктов
	if err := productsQuery.Limit(pageSize).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	// Затем загружаем ассоциации для каждого продукта
	for i := range products {
		manufacturerQuery := repo.db.Model(&types.Manufacturer{})
		if err := manufacturerQuery.First(&products[i].Manufacturer, products[i].ManufacturerID).Error; err != nil {
			return nil, 0, err
		}
		categoryQuery := repo.db.Model(&types.Category{})
		if err := categoryQuery.Select("category.*").
			Joins("JOIN product_category ON product_category.category_id = category.category_id").
			Joins("JOIN product ON product.product_id = product_category.product_id").
			Find(&products[i].Category, products[i].ProductID).Error; err != nil {
			return nil, 0, err
		}
		characteristicQuery := repo.db.Model(&types.Characteristic{})
		if err := characteristicQuery.Select("characteristic.*, product_characteristic.value").
			Joins("JOIN product_characteristic ON "+
				"product_characteristic.characteristic_id = characteristic.characteristic_id").
			Joins("JOIN product ON product.product_id = product_characteristic.product_id").
			Find(&products[i].Characteristic, products[i].ProductID).Error; err != nil {
			return nil, 0, err
		}
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

package data

import (
	"errors"
	"gorm.io/gorm"
	"online-shop-API/internal/types"
	_ "strconv"
)

// NewProductRepository creates a new instance of Repository.
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// GetProducts retrieves a list of products with optional filters and pagination.
func (repo *Repository) GetProducts(
	filters map[string]interface{},
	page int,
	pageSize int,
) ([]types.Product, int64, error) {

	repo.db = repo.db.Session(&gorm.Session{NewDB: true}).Debug()
	repo.db = repo.db.Set("gorm:auto_preload", true)
	var products []types.Product
	var totalCount int64 = 0

	// Initialize the productsQuery
	productsQuery := repo.db.Model(&types.Product{})

	// Apply filters
	for key, value := range filters {
		productsQuery = productsQuery.Where(key+" = ?", value)
	}

	// Count total products for pagination
	if err := productsQuery.
		Joins("JOIN product_category ON product.product_id = product_category.product_id").
		Joins("JOIN category ON product_category.category_id = category.category_id").
		Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	// Сначала получаем список продуктов
	if err := productsQuery.
		Preload("Manufacturer").
		Preload("Category.Category").
		Preload("Characteristic.Characteristic").
		Limit(pageSize).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

// AddProduct добавляет новый продукт в базу данных.
func (repo *Repository) AddProduct(product *types.Product) (*types.Product, error) {
	// Проверка на обязательные поля
	if product.Name == "" {
		return nil, errors.New("product name is required")
	}

	//Если статус не установлен, устанавливается в активный
	if product.Status == "" {
		product.Status = "ACTIVE"
	}

	// Сохраняем продукт в базе
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}

	err := repo.SubscribeProduct(product.ProductID, 0)
	if err != nil {
		return nil, err
	}

	// Возвращаем созданный продукт
	return product, nil
}

// UpdateProduct обновляет существующий продукт в базе данных.
func (repo *Repository) UpdateProduct(
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

	if updatedData.Category != nil {
		product.Category = updatedData.Category
	}

	if updatedData.Characteristic != nil {
		product.Characteristic = updatedData.Characteristic
	}

	// Сохраняем изменения
	if err := repo.db.Save(&product).Error; err != nil {
		return nil, err
	}

	// Возвращаем обновленный продукт
	return &product, nil
}

// DeleteProduct удаляет продукт из базы данных по его ID.
func (repo *Repository) DeleteProduct(productID uint) error {
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

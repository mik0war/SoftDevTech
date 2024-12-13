package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"online-shop-API/internal/types"
	"os"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "root", "online-shop")

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграция схемы
	err = Db.AutoMigrate(&types.Product{})
	if err != nil {
		return nil
	}
	err = Db.AutoMigrate(&types.Category{})
	if err != nil {
		return nil
	}
	err = Db.AutoMigrate(&types.Characteristic{})
	if err != nil {
		return nil
	}
	err = Db.AutoMigrate(&types.User{})
	if err != nil {
		return nil
	}
	err = Db.AutoMigrate(&types.ProductCategory{})
	if err != nil {
		return nil
	}

	return Db
}

// Repository provides access to the product store.
type Repository struct {
	db *gorm.DB
}

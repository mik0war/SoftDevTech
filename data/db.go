package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"online-shop-API/types"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=localhost " +
		"user=postgres password=root dbname=online-shop port=5432 sslmode=disable"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
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
	err = Db.AutoMigrate(&types.User{})
	if err != nil {
		return nil
	}

	return Db
}
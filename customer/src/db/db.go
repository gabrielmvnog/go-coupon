package db

import (
	"log"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "postgres://customer:customer@0.0.0.0:5432/customer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(models.Customer{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}

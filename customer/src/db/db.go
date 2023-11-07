package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(models.Customer{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}

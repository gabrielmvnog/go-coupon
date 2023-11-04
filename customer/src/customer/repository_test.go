package customer

import (
	"context"
	"log"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetCustomerById(t *testing.T) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	customerRepo := NewCustomerRepository(db)

	customer, _ := customerRepo.GetCustomerById(context.TODO(), 123)

	log.Print(customer)
}

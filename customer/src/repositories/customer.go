package repositories

import (
	"context"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r CustomerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	result := r.db.Create(customer)

	return customer, result.Error
}

func (r CustomerRepository) GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error) {
	var customer models.Customer
	result := r.db.First(&customer, customer_id)

	return &customer, result.Error
}

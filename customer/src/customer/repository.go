package customer

import (
	"context"

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

func (r CustomerRepository) CreateCustomer(ctx context.Context, customer *Customer) (*Customer, error) {
	result := r.db.Create(customer)

	return customer, result.Error
}

func (r CustomerRepository) GetCustomerById(ctx context.Context, customer_id uint32) (*Customer, error) {
	var customer Customer
	result := r.db.First(&customer, customer_id)

	return &customer, result.Error
}

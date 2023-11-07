package repositories

import (
	"context"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer_id uint32, customer *models.Customer) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r customerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	result := r.db.Create(customer)

	return customer, result.Error
}

func (r customerRepository) GetCustomerById(ctx context.Context, customer_id uint32) (*models.Customer, error) {
	var customer models.Customer
	result := r.db.First(&customer, customer_id)

	return &customer, result.Error
}

func (r customerRepository) UpdateCustomer(ctx context.Context, customer_id uint32, customer *models.Customer) (*models.Customer, error) {
	result := r.db.Model(models.Customer{}).Where("id = ?", customer_id).Updates(customer)

	return customer, result.Error
}

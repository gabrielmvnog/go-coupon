package repositories_test

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/gabrielmvnog/go-coupon/customer/src/models"
	"github.com/gabrielmvnog/go-coupon/customer/src/repositories"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetCustomerById(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	customerRepo := repositories.NewCustomerRepository(db)

	mock.ExpectQuery(
		regexp.QuoteMeta(`SELECT * FROM "customers" WHERE "customers"."id" = $1 ORDER BY "customers"."id" LIMIT 1`),
	).WithArgs(uint32(1)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	customer, err := customerRepo.GetCustomerById(context.Background(), uint32(1))

	assert.Equal(t, uint32(1), customer.ID, "Customer should have id in return")
	assert.Equal(t, nil, err)
}

func TestCreateCustomer(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	customerRepo := repositories.NewCustomerRepository(db)

	mock.ExpectBegin()
	mock.ExpectQuery(
		regexp.QuoteMeta(`INSERT INTO "customers" ("first_name","last_name","email","phone","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`),
	).WithArgs("", "", "", "", AnyTime{}, AnyTime{}).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	customer, err := customerRepo.CreateCustomer(context.Background(), &models.Customer{})

	assert.Equal(t, uint32(1), customer.ID, "Customer should have id in return")
	assert.Equal(t, nil, err)
}

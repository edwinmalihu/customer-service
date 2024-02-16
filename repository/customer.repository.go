package repository

import (
	"customer-service/model"
	"customer-service/request"
	"log"

	"gorm.io/gorm"
)

type CustomerRepo interface {
	Migrate() error
	AddCustomer(request.AddCustomer) (model.Customer, error)
}

type customerRepo struct {
	DB *gorm.DB
}

// AddCustomer implements CustomerRepo.
func (customerRepo) AddCustomer(request.AddCustomer) (model.Customer, error) {
	panic("unimplemented")
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	return customerRepo{
		DB: db,
	}
}

func (u customerRepo) Migrate() error {
	log.Print("[CustomerRepository]...Migrate")
	return u.DB.AutoMigrate(&model.Customer{})
}

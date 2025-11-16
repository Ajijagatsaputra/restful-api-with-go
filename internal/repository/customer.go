package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"shellrean.id/belajar-golang-rest-api/domain"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements domain.CustomerRepository.
func (cr *customerRepository) FindAll(ctx *context.Context) ([]domain.Customer, error) {
	panic("implement me")
}

// FindByID implements domain.CustomerRepository.
func (cr *customerRepository) FindByID(ctx *context.Context, id string) (domain.Customer, error) {
	panic("implement me")
}

// Save implements domain.CustomerRepository.
func (cr *customerRepository) Save(ctx *context.Context, customer domain.Customer) error {
	panic("implement me")
}

// Update implements domain.CustomerRepository.
func (cr *customerRepository) Update(ctx *context.Context, customer domain.Customer) error {
	panic("implement me")
}

// Delete implements domain.CustomerRepository.
func (cr *customerRepository) Delete(ctx *context.Context, id string) error {
	panic("implement me")
}

package domain

import (
	"context"
	"database/sql"
)

type Customer struct {
	ID         string       `db:"id"`
	Code       string       `db:"code"`
	Name       string       `db:"name"`
	CreatedAt  sql.NullTime `db:"created_at"`
	UppdatedAt sql.NullTime `db:"updated_at"`
	DeleteAt   sql.NullTime `db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx *context.Context) ([]Customer, error)
	FindByID(ctx *context.Context, id string) (Customer, error)
	Save(ctx *context.Context, customer Customer) error
	Update(ctx *context.Context, customer Customer) error
	Delete(ctx *context.Context, id string) error
}

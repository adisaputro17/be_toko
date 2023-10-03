package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_toko/entity"
)

type DomainItf interface {
	// Account
	CreateAccount(ctx context.Context, v entity.Account) (entity.Account, error)
	GetAccountByUsernameAndPassword(ctx context.Context, p entity.LoginRequest) (entity.Account, error)

	// Product
	GetProductCategory(ctx context.Context) ([]entity.ProductCategory, error)
	GetProductByCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error)
	GetProductByProductID(ctx context.Context, productID string) (entity.Product, error)

	// Cart
	InsertCart(ctx context.Context, v entity.Cart) (entity.Cart, error)
}

type domain struct {
	DB  *sql.DB
	Opt Options
}

type Options struct {
}

func InitDomain(
	db *sql.DB,
	opt Options,
) DomainItf {
	return &domain{
		DB:  db,
		Opt: opt,
	}
}

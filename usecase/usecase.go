package usecase

import (
	"context"

	"github.com/adisaputro17/be_toko/domain"
	"github.com/adisaputro17/be_toko/entity"
	"github.com/go-playground/validator/v10"
)

type UsecaseItf interface {
	// Account
	CreateAccount(ctx context.Context, request entity.CreateAccountRequest) (entity.Account, error)
	Login(ctx context.Context, request entity.LoginRequest) (entity.Account, error)

	// Product
	GetProductCategory(ctx context.Context) ([]entity.ProductCategory, error)
	GetProductByCategoryID(ctx context.Context, categoryID string) ([]entity.Product, error)
	GetProductByProductID(ctx context.Context, productID string) (entity.Product, error)

	// Cart
	InsertCart(ctx context.Context, request entity.InsertCartRequest) (entity.Cart, error)
}

type usecase struct {
	Domain   domain.DomainItf
	Validate *validator.Validate
}

func InitUsecase(domain domain.DomainItf, validate *validator.Validate) UsecaseItf {
	return &usecase{
		Domain:   domain,
		Validate: validate,
	}
}

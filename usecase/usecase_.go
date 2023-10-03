package usecase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/adisaputro17/be_toko/entity"
)

func (u *usecase) CreateAccount(ctx context.Context, request entity.CreateAccountRequest) (entity.Account, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Account{}, err
	}

	timeNowUTC := time.Now().UTC()
	accountID := fmt.Sprint(timeNowUTC.Unix())
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}
	accountRequest := entity.Account{
		AccountID:   accountID,
		Fullname:    request.Fullname,
		Username:    request.Username,
		Password:    base64.StdEncoding.EncodeToString([]byte(request.Password)),
		CommonField: commonField,
	}

	account, err := u.Domain.CreateAccount(ctx, accountRequest)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (u *usecase) Login(ctx context.Context, request entity.LoginRequest) (entity.Account, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Account{}, err
	}

	request.Password = base64.StdEncoding.EncodeToString([]byte(request.Password))
	account, err := u.Domain.GetAccountByUsernameAndPassword(ctx, request)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (u *usecase) GetProductCategory(ctx context.Context) ([]entity.ProductCategory, error) {
	return u.Domain.GetProductCategory(ctx)
}

func (u *usecase) GetProductByCategoryID(ctx context.Context, categoryID string) ([]entity.Product, error) {
	return u.Domain.GetProductByCategoryID(ctx, categoryID)
}

func (u *usecase) GetProductByProductID(ctx context.Context, productID string) (entity.Product, error) {
	return u.Domain.GetProductByProductID(ctx, productID)
}

func (u *usecase) InsertCart(ctx context.Context, request entity.InsertCartRequest) (entity.Cart, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Cart{}, err
	}

	product, err := u.GetProductByProductID(ctx, request.ProductID)
	if err != nil {
		return entity.Cart{}, err
	}

	if request.Qty > product.Stock {
		return entity.Cart{}, errors.New("quantity exceeds limit")
	}

	timeNowUTC := time.Now().UTC()
	cartID := fmt.Sprintf("%s-%s", request.AccountID, request.ProductID)
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}
	cartRequest := entity.Cart{
		CartID:      cartID,
		AccountID:   request.AccountID,
		ProductID:   request.ProductID,
		Qty:         request.Qty,
		CommonField: commonField,
	}

	cart, err := u.Domain.InsertCart(ctx, cartRequest)
	if err != nil {
		return entity.Cart{}, err
	}

	return cart, nil
}

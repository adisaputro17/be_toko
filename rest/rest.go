package rest

import (
	"github.com/adisaputro17/be_toko/usecase"
	"github.com/labstack/echo/v4"
)

type REST interface {
	CreateAccount(c echo.Context) error
	Login(c echo.Context) error

	GetProductCategory(c echo.Context) error
	GetProductByCategoryID(c echo.Context) error
	GetProductByProductID(c echo.Context) error

	InsertCart(c echo.Context) error
}

type rest struct {
	Usecase usecase.UsecaseItf
}

func Init(usecase usecase.UsecaseItf) REST {
	return &rest{
		Usecase: usecase,
	}
}

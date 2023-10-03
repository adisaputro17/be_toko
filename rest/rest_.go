package rest

import (
	"net/http"
	"time"

	"github.com/adisaputro17/be_toko/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (rest *rest) CreateAccount(c echo.Context) error {
	req := new(entity.CreateAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	account, err := rest.Usecase.CreateAccount(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.HTTPResp{
		Data: account,
	})
}

func (rest *rest) Login(c echo.Context) error {
	req := new(entity.LoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	account, err := rest.Usecase.Login(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	// Set custom claims
	claims := &entity.JWTCustomClaims{
		account.AccountID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (rest *rest) GetProductCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*entity.JWTCustomClaims)

	productCategory, err := rest.Usecase.GetProductCategory(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: productCategory,
	})
}

func (rest *rest) GetProductByCategoryID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*entity.JWTCustomClaims)

	categoryID := c.Param("category_id")

	products, err := rest.Usecase.GetProductByCategoryID(c.Request().Context(), categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: products,
	})
}

func (rest *rest) GetProductByProductID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_ = user.Claims.(*entity.JWTCustomClaims)

	productID := c.Param("product_id")

	product, err := rest.Usecase.GetProductByProductID(c.Request().Context(), productID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: product,
	})
}

func (rest *rest) InsertCart(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JWTCustomClaims)

	req := new(entity.InsertCartRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	req.AccountID = claims.AccountID

	cart, err := rest.Usecase.InsertCart(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.HTTPResp{
		Data: cart,
	})
}

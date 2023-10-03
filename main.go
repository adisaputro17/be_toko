package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adisaputro17/be_toko/app"
	"github.com/adisaputro17/be_toko/domain"
	"github.com/adisaputro17/be_toko/entity"
	"github.com/adisaputro17/be_toko/rest"
	"github.com/adisaputro17/be_toko/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := entity.AppConfig{
		AppPort: os.Getenv("APP_PORT"),
		DBConfig: entity.DatabaseConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBName:     os.Getenv("DB_NAME"),
		},
	}

	validate := validator.New()
	db, err := app.NewDB(appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	opt := domain.Options{}

	domain := domain.InitDomain(db, opt)
	usecase := usecase.InitUsecase(domain, validate)
	rest := rest.Init(usecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/daftar", rest.CreateAccount)
	e.POST("/login", rest.Login)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTCustomClaims)
		},
		SigningKey: []byte(entity.Sign),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("/category", rest.GetProductCategory)
	r.GET("/product/category/:category_id", rest.GetProductByCategoryID)
	r.GET("/product/:product_id", rest.GetProductByProductID)
	r.POST("/cart", rest.InsertCart)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appConfig.AppPort)))

}

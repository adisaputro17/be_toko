package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	Sign = "adhisaputro17"
)

type AppConfig struct {
	AppPort  string
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

type CommonField struct {
	CreatedAt time.Time `json:"-"`
	CreatedBy int64     `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UpdatedBy int64     `json:"-"`
}

type JWTCustomClaims struct {
	AccountID string `json:"account_id"`
	jwt.RegisteredClaims
}

type HTTPResp struct {
	Data interface{} `json:"data,omitempty"`
}

type HTTPErrorResponse struct {
	Message string `json:"message"`
}

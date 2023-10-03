package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
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

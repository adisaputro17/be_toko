package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adisaputro17/be_toko/entity"
)

func NewDB(dbConfig entity.DatabaseConfig) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return &sql.DB{}, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}

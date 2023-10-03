package domain

import (
	"database/sql"

	"github.com/adisaputro17/be_toko/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlInsertCart(tx *sql.Tx, v entity.Cart) (*sql.Tx, error) {
	_, err := tx.Exec(insertCart,
		v.CartID,
		v.AccountID,
		v.ProductID,
		v.Qty,
		v.CreatedAt,
		v.CreatedBy,
		v.UpdatedAt,
		v.UpdatedBy,
	)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

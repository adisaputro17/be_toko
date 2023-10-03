package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_toko/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlCreateAccount(tx *sql.Tx, v entity.Account) (*sql.Tx, error) {
	_, err := tx.Exec(createAccount,
		v.AccountID,
		v.Fullname,
		v.Username,
		v.Password,
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

func (d *domain) sqlReadAccountByUsernameAndPassword(ctx context.Context, p entity.LoginRequest) (entity.Account, error) {
	result := entity.Account{}
	row := d.DB.QueryRowContext(ctx, readAccountByUsernameAndPassword, p.Username, p.Password)

	err := row.Scan(
		&result.AccountID,
		&result.Fullname,
		&result.Username,
		&result.Password,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

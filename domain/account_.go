package domain

import (
	"context"
	"database/sql"
	"log"

	"github.com/adisaputro17/be_toko/entity"
)

func (d *domain) CreateAccount(ctx context.Context, v entity.Account) (entity.Account, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlCreateAccount(tx, v)
	if err != nil {
		return v, err
	}

	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			tx.Commit()
		}
	}()

	return v, nil
}

func (d *domain) GetAccountByUsernameAndPassword(ctx context.Context, p entity.LoginRequest) (entity.Account, error) {
	return d.sqlReadAccountByUsernameAndPassword(ctx, p)
}

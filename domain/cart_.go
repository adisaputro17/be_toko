package domain

import (
	"context"
	"database/sql"
	"log"

	"github.com/adisaputro17/be_toko/entity"
)

func (d *domain) InsertCart(ctx context.Context, v entity.Cart) (entity.Cart, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlInsertCart(tx, v)
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

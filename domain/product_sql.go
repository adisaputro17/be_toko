package domain

import (
	"context"

	"github.com/adisaputro17/be_toko/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlReadProductCategory(ctx context.Context) ([]entity.ProductCategory, error) {
	results := []entity.ProductCategory{}
	rows, err := d.DB.QueryContext(ctx, readProductCategory)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.ProductCategory{}
		if err := rows.Scan(
			&data.ProductCategoryID,
			&data.ProductCategoryName,
			&data.CreatedAt,
			&data.CreatedBy,
			&data.UpdatedAt,
			&data.UpdatedBy,
		); err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}

func (d *domain) sqlReadProductByCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error) {
	results := []entity.Product{}
	rows, err := d.DB.QueryContext(ctx, readProductByCategoryID, productCategoryID)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Product{}
		if err := rows.Scan(
			&data.ProductID,
			&data.ProductName,
			&data.Category.ProductCategoryID,
			&data.Category.ProductCategoryName,
			&data.Description,
			&data.Price,
			&data.Stock,
			&data.CreatedAt,
			&data.CreatedBy,
			&data.UpdatedAt,
			&data.UpdatedBy,
		); err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}

func (d *domain) sqlReadProductByProductID(ctx context.Context, productID string) (entity.Product, error) {
	result := entity.Product{}
	row := d.DB.QueryRowContext(ctx, readProductByProductID, productID)

	err := row.Scan(
		&result.ProductID,
		&result.ProductName,
		&result.Category.ProductCategoryID,
		&result.Category.ProductCategoryName,
		&result.Description,
		&result.Price,
		&result.Stock,
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

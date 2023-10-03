package domain

import (
	"context"

	"github.com/adisaputro17/be_toko/entity"
)

func (d *domain) GetProductCategory(ctx context.Context) ([]entity.ProductCategory, error) {
	return d.sqlReadProductCategory(ctx)
}

func (d *domain) GetProductByCategoryID(ctx context.Context, productCategoryID string) ([]entity.Product, error) {
	return d.sqlReadProductByCategoryID(ctx, productCategoryID)
}

func (d *domain) GetProductByProductID(ctx context.Context, productID string) (entity.Product, error) {
	return d.sqlReadProductByProductID(ctx, productID)
}

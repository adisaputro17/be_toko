package domain

const (
	readProductCategory = `SELECT
		product_category_id,
		product_category_name,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM product_category`

	readProductByCategoryID = `SELECT
		p.product_id,
		p.product_name,
		pc.product_category_id,
		pc.product_category_id,
		p.description,
		p.price,
		p.stock,
		p.created_at,
		p.created_by,
		p.updated_at,
		p.updated_by
	FROM product p
	JOIN product_category pc ON p.product_category_id = pc.product_category_id
	WHERE p.product_category_id = ?`

	readProductByProductID = `SELECT
		p.product_id,
		p.product_name,
		pc.product_category_id,
		pc.product_category_id,
		p.description,
		p.price,
		p.stock,
		p.created_at,
		p.created_by,
		p.updated_at,
		p.updated_by
	FROM product p
	JOIN product_category pc ON p.product_category_id = pc.product_category_id
	WHERE p.product_id = ?`
)

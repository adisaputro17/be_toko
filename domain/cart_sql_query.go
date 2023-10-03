package domain

const (
	insertCart = `INSERT INTO cart (
		cart_id,
		account_id,
		product_id,
		qty,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
)

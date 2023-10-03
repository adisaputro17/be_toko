package domain

const (
	createAccount = `INSERT INTO account (
		account_id,
		fullname,
		username,
		password,
		created_at,
		created_by,
		updated_at,
		updated_by
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	readAccountByUsernameAndPassword = `SELECT
		account_id,
		fullname,
		username,
		password,
		created_at,
		created_by,
		updated_at,
		updated_by
	FROM account WHERE username = ? AND password = ?`
)

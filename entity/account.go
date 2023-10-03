package entity

type Account struct {
	AccountID string `json:"account_id"`
	Fullname  string `json:"fullname"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CommonField
}

type CreateAccountRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

package entity

type Cart struct {
	CartID    string `json:"cart_id"`
	AccountID string `json:"account_id"`
	ProductID string `json:"product_id"`
	Qty       int64  `json:"qty"`
	CommonField
}

type InsertCartRequest struct {
	AccountID string `json:"account_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Qty       int64  `json:"qty" validate:"required"`
}

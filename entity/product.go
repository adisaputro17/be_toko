package entity

type ProductCategory struct {
	ProductCategoryID   string `json:"product_category_id"`
	ProductCategoryName string `json:"product_category_name"`
	CommonField
}

type Product struct {
	ProductID   string          `json:"product_id"`
	ProductName string          `json:"product_name"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Stock       int64           `json:"stock"`
	Category    ProductCategory `json:"category"`
	CommonField
}

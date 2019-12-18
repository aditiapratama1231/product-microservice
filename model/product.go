package model

type Product struct {
	ProductName string `json:"product_name" gorm:"column:product_name"`
	SKU         string `json:"sku" gorm:"column:sku"`
	Qty         int32  `json:"qty" gorm:"column:qty"`
}

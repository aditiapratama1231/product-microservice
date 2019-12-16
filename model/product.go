package model

type Product struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	ProductName string `json:"product_name" gorm:"column:product_name"`
	SKU         string `json:"sku" gorm:"column:sku"`
	Qty         int32  `json:"qty" gorm:"column:qty"`
}

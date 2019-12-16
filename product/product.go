package product

import (
	"library/model"

	"github.com/jinzhu/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) Product {
	return Product{
		DB: db,
	}
}

func (p *Product) GetProducts() []model.Product {
	db := p.DB
	var products []model.Product

	db.Find(&products)

	return products
}

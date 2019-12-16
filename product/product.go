package product

import (
	"library/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Product struct {
	DB *gorm.DB
}

type product struct {
	ProductName string `json:"product_name"`
	SKU         string `json:"sku"`
	Qty         int32  `json:"qty"`
}

func (p *Product) GetProducts(c *gin.Context) {
	db := p.DB
	var products []model.Product

	db.Find(&products)

	c.JSON(200, gin.H{
		"data": products,
	})
}

func (p *Product) CreateProduct(c *gin.Context) {
	var request product

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	p.DB.Create(&request)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

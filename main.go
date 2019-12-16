package main

import (
	"github.com/gin-gonic/gin"
	"library/config"
	"library/product"
)

func main() {
	db := config.DBInit()
	prod := product.NewProduct(db)

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": prod.GetProducts(),
		})
	})

	r.Run()
}

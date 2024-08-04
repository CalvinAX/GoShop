package controllers

import (
	"net/http"

	"github.com/CalvinAX/GoShop/initializers"
	"github.com/CalvinAX/GoShop/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var body struct {
		Name          string
		Price         float64
		ProductNumber string
		Stock         int
		TaxID         int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body on CreateProduct request!",
		})

		return
	}

	product := models.Product{Name: body.Name, Price: body.Price, ProductNumber: body.ProductNumber, Stock: body.Stock, TaxID: body.TaxID}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create product!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product successfully created!",
		"product": product,
	})
}

package controllers

import (
	"golang-mnc/initializer"
	"golang-mnc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpMerchant(c *gin.Context) {
	var body struct {
		Name    string
		Address string
		Phone   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Read Body",
		})
		return
	}

	merchant := models.Merchant{Name: body.Name, Address: body.Address, Phone: body.Phone}
	result := initializer.DB.Create(&merchant)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Save Merchant",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

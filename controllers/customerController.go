package controllers

import (
	"golang-mnc/initializer"
	"golang-mnc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpCustomer(c *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		Address   string
		Phone     string
		User      models.User
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.User.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	customer := models.Customer{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Address:   body.Address,
		Phone:     body.Phone,
		User:      models.User{Email: body.User.Email, Password: string(hash)},
	}
	result := initializer.DB.Create(&customer)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Registered",
	})

}

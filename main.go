package main

import (
	"golang-mnc/controllers"
	"golang-mnc/initializer"
	"golang-mnc/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDatabase()
	initializer.SyncDatbase()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.Login)
	r.POST("/register-merchant", controllers.SignUpMerchant)
	r.POST("/register-customer", controllers.SignUpCustomer)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}

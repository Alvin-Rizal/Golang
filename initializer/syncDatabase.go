package initializer

import "golang-mnc/models"

func SyncDatbase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Merchant{})
	DB.AutoMigrate(&models.Customer{})
}

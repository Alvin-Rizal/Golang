package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string `gorm:"unique"`
	Phone     string `gorm:"unique"`
	UserId    int
	User      User `gorm:"references:id"`
}

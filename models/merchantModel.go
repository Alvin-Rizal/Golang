package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name    string
	Address string `gorm:"unique"`
	Phone   string `gorm:"unique"`
}

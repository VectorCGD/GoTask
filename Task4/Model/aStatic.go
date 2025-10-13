package model

import (
	base "BockWeb/Database"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = base.Database

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

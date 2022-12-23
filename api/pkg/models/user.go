package models

import (
	"api/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}
func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

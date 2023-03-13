package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Password   string
	AdminPower bool `gorm:"default:false"`
}

type UserInfo struct {
	gorm.Model
	User_ID   uint
	User      User   `gorm:"foreignKey:User_ID"`
	Email     string `gorm:"unique"`
	FirstName string
	LastName  string
}

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
	ID        int `gorm:"primary_key;auto_increment;not_null"`
	UserRefer int
	User      User   `gorm:"foreignKey:UserRefer"`
	Email     string `gorm:"unique"`
	FirstName string
	LastName  string
}

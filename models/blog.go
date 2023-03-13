package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	User_ID  int
	User     User `gorm:"foreignKey:User_ID"`
	Content  string
	Comments []Comment `gorm:"many2many:user_comment;"`
}

type Comment struct {
	gorm.Model
	User_ID        int
	User           User `gorm:"foreignKey:User_ID"`
	CommentContent string
}

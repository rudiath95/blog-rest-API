package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	User_ID  uint
	User     User `gorm:"foreignKey:User_ID"`
	Content  string
	Comments []Comment `gorm:"foreignKey:CommentRefer"` //`gorm:"many2many:user_comment;"`
}

type Comment struct {
	gorm.Model
	User_ID        int
	User           User `gorm:"foreignKey:User_ID"`
	CommentRefer   uint
	CommentContent string
}

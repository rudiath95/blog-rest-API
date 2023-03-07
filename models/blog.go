package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	User_ID        int
	OriginalPoster User `gorm:"foreignKey:User_ID"`
	Content        string
}

type Comment struct {
	gorm.Model
	Blog_ID        int
	Blog           Blog `gorm:"foreignKey:Blog_ID"`
	User_ID        int
	User           User `gorm:"foreignKey:Blog_ID"`
	CommentContent string
}

package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	OriginalPoster string
	Username       User `gorm:"foreignKey:OriginalPoster"`
	Content        string
}

type Comment struct {
	gorm.Model
	Blog_ID        int
	Blog           Blog `gorm:"foreignKey:Blog_ID"`
	Commenter      string
	Username       User `gorm:"foreignKey:Commenter"`
	CommentContent string
}

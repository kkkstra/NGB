package model

import (
	"gorm.io/gorm"
)

// 点赞表
type Thumbs struct {
	gorm.Model
	UserID uint `gorm:"primaryKey"`
	PostID uint `gorm:"primaryKey"`
}

package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	UserID     uint
	CategoryID uint
}

package model

import (
	"NGB/pkg/logrus"
	"time"

	"gorm.io/gorm"
)

// 点赞表
type UserThumbs struct {
  UserID	uint	`gorm:"primaryKey"`
  PostID	uint	`gorm:"primaryKey"`
  CreatedAt	time.Time
  DeletedAt	gorm.DeletedAt
}

func init() {
	err := db.SetupJoinTable(&User{}, "Thumbs", &UserThumbs{})
	if err != nil {
		logrus.Logger.Error(err)
	}
}
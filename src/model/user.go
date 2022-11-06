package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Intro    string
}

func (model *model) Insert(u *User) (uint, error) {
	res := model.db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, res.Error
}

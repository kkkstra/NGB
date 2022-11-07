package model

import (
	"gorm.io/gorm"
)

type RoleType int

type UserModel struct {
	gorm.Model
	Username string   `gorm:"primarykey;unique;not null"`
	Email    string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Role     RoleType `gorm:"not null;default:0"`
	Intro    string
	Github   string
	School   string
	Website  string
}

const (
	common RoleType = 0
	admin  RoleType = 1
)

func (m *model) CreateUser(u *UserModel) (uint, error) {
	res := m.db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, res.Error
}

func (m *model) FindUser(username string) (*UserModel, error) {
	res := &UserModel{}
	tx := db.First(res,
		"username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

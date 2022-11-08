package model

import (
	"gorm.io/gorm"
)

type RoleType int

type User struct {
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

type UserModel struct {
	db *gorm.DB
}

const (
	common RoleType = 0
	admin  RoleType = 1
)

func (m *UserModel) CreateUser(u *User) (uint, error) {
	res := m.db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, res.Error
}

func (m *UserModel) FindUser(username string) (*User, error) {
	res := &User{}
	tx := db.First(res,
		"username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (r *RoleType) Str() string {
	switch *r {
	case common:
		return "common"
	case admin:
		return "admin"
	}
	return "others"
}

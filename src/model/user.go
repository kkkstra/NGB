package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"primarykey;unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     bool   `gorm:"not null;default:false"` // true: admin, false: common user
	Intro    string
	Github   string
	School   string
	Website  string
}

func (m *model) CreateUser(u *User) (uint, error) {
	res := m.db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, res.Error
}

func (m *model) FindUser(username string) (*User, error) {
	res := &User{}
	tx := db.First(res, "username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

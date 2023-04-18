package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string `gorm:"primarykey;not null"`
	Email            string `gorm:"not null"`
	Password         string `gorm:"not null"`
	Role             int    `gorm:"not null;default:0"`
	Intro            string
	UpdatePasswordAt time.Time
}

func (m *Model) CreateUser(u *User) (uint, error) {
	res := m.db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, nil
}

func (m *Model) FindUserByUsername(username string) (*User, error) {
	res := &User{}
	tx := db.First(res,
		"username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (m *Model) FindUserById(id string) (*User, error) {
	var user User
	tx := db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (m *Model) UpdateUser(id string, u *User) error {
	var user User
	db.First(&user, id)
	tx := db.Model(&user).Updates(u)
	return tx.Error
}

func (m *Model) DelUser(id string) error {
	tx := db.Delete(&User{}, id)
	return tx.Error
}

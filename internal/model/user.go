package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string `gorm:"not null"`
	Email            string `gorm:"not null"`
	Password         string `gorm:"not null"`
	Role             int    `gorm:"not null;default:0"`
	Intro            string
	UpdatePasswordAt time.Time
}

type Following struct {
	gorm.Model
	UserID      uint `gorm:"primaryKey"`
	FollowingID uint `gorm:"primaryKey"`
}

const (
	roleCommon = 0
	roleAdmin  = 1
)

func (m *Model) CreateUser(u *User) (uint, error) {
	tx := m.db.Create(u)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return u.ID, nil
}

func (m *Model) FindUserByUsername(username string) (*User, error) {
	res := &User{}
	tx := m.db.First(res,
		"username = ?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (m *Model) FindUserByEmail(email string) (*User, error) {
	res := &User{}
	tx := m.db.First(res,
		"email = ?", email)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (m *Model) FindUserById(id string) (*User, error) {
	var user User
	tx := m.db.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (m *Model) UpdateUser(id string, u *User) error {
	var user User
	m.db.First(&user, id)
	tx := m.db.Model(&user).Updates(u)
	return tx.Error
}

func (m *Model) DelUser(id string) error {
	tx := m.db.Delete(&User{}, id)
	return tx.Error
}

func (m *Model) CreateFollowing(f *Following) error {
	tx := m.db.Create(f)
	return tx.Error
}

// get record id
func (m *Model) GetUserFollowingID(f *Following) (uint, error) {
	res := &Following{}
	tx := m.db.First(res,
		"user_id = ? AND following_id = ?", f.UserID, f.FollowingID)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return res.ID, nil
}

func (m *Model) DeleteFollowing(f *Following) error {
	id, err := m.GetUserFollowingID(f)
	if err != nil {
		return err
	}
	tx := m.db.Delete(&Following{}, id)
	return tx.Error
}

func (m *Model) GetAllFollowings(userID string) ([]uint, error) {
	res := []Following{}
	tx := m.db.Where("user_id = ?", userID).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	followings := []uint{}
	for _, f := range res {
		followings = append(followings, f.FollowingID)
	}
	return followings, nil
}

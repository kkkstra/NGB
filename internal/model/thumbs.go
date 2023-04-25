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

func (m *Model) GetThumbsByPost(postID string) ([]uint, error) {
	res := []Thumbs{}
	tx := m.db.Where("post_id = ?", postID).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	users := []uint{}
	for _, t := range res {
		users = append(users, t.UserID)
	}
	return users, nil
}

func (m *Model) GetThumbsByUser(userID string) ([]uint, error) {
	res := []Thumbs{}
	tx := m.db.Where("user_id = ?", userID).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	posts := []uint{}
	for _, t := range res {
		posts = append(posts, t.PostID)
	}
	return posts, nil
}

func (m *Model) CreateThumbs(t *Thumbs) (uint, error) {
	tx := m.db.Create(t)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return t.ID, nil
}

// get record id
func (m *Model) GetThumbsID(t *Thumbs) (uint, error) {
	res := &Thumbs{}
	tx := m.db.First(res,
		"user_id = ? AND post_id = ?", t.UserID, t.PostID)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return res.ID, nil
}

func (m *Model) DeleteThumbs(t *Thumbs) error {
	id, err := m.GetThumbsID(t)
	if err != nil {
		return err
	}
	tx := m.db.Delete(&Thumbs{}, id)
	return tx.Error
}

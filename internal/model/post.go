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

func (m *Model) GetPost(id string) (*Post, error) {
	var post Post
	tx := m.db.First(&post, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &post, nil
}

func (m *Model) GetPostsByCategory(id string) {
}

func (m *Model) GetPostsByUser(userID string) ([]Post, error) {
	posts := []Post{}
	tx := m.db.Where("user_id = ?", userID).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}

func (m *Model) CreatePost(p *Post) (uint, error) {
	tx := m.db.Create(p)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return p.ID, nil
}

func (m *Model) DelPost(id string) error {
	tx := m.db.Delete(&Post{}, id)
	return tx.Error
}

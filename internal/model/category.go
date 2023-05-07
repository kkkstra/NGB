package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (m *Model) GetAllCategories() ([]Category, error) {
	categories := []Category{}
	tx := m.db.Find(&categories)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categories, nil
}

func (m *Model) GetCategory(id string) (*Category, error) {
	var category Category
	tx := m.db.First(&category, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}

func (m *Model) CreateCategory(p *Category) (uint, error) {
	tx := m.db.Create(p)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return p.ID, nil
}

func (m *Model) DelCategory(id string) error {
	tx := m.db.Delete(&Category{}, id)
	return tx.Error
}

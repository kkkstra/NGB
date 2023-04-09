package model

import (
	"gorm.io/gorm"
)

type Model struct {
	db *gorm.DB
}

// TODO
// 这里有无更好的写法？
func GetModel() *Model {
	return &Model{db}
}

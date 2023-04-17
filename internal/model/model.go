package model

import (
	"gorm.io/gorm"
)

type Model struct {
	db *gorm.DB
}

func GetModel() Model {
	return Model{db}
}

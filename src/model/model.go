package model

import (
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func GetModel() *model {
	return &model{db}
}

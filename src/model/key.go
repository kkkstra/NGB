package model

import (
	"byitter/src/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RSAKeyModel struct {
	gorm.Model
	Kid        string
	Type       string
	PublicKey  string
	PrivateKey string
}

func (m *model) CreateRSAKey(keyType string) {
	publicKey, privateKey := config.ReadRSAKeyFromFile(keyType)
	key := RSAKeyModel{
		Kid:        uuid.New().String(),
		Type:       keyType,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
	res := m.db.Create(&key)
	if res.Error != nil {
		panic(res.Error)
	}
}

func (m *model) FindRSAKey() []RSAKeyModel {
	var keyList []RSAKeyModel
	res := m.db.Find(&keyList)
	if res.Error != nil {
		return nil
	}
	return keyList
}

package model

import (
	"byitter/src/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RSAKey struct {
	gorm.Model
	Kid        string
	Type       string
	PublicKey  string
	PrivateKey string
}

type RSAKeyModel struct {
	db *gorm.DB
}

func (m *RSAKeyModel) CreateRSAKey(keyType string) {
	publicKey, privateKey := config.ReadRSAKeyFromFile(keyType)
	key := RSAKey{
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

func (m *RSAKeyModel) FindRSAKey() []RSAKey {
	var keyList []RSAKey
	res := m.db.Find(&keyList)
	if res.Error != nil {
		return nil
	}
	return keyList
}

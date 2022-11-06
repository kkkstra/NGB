package model

import (
	"byitter/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Intro    string
}

var db *gorm.DB

// ConnectDatabase 连接数据库
func ConnectDatabase() {
	dsn := "host=" + config.C.Postgresql.Host + " user=" + config.C.Postgresql.User + " password=" + config.C.Postgresql.Password + " dbname=" +
		config.C.Postgresql.Dbname + " port=" + config.C.Postgresql.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// MigrateSchema 迁移schema
func MigrateSchema() {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func (u *User) Insert() (uint, error) {
	res := db.Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, res.Error
}

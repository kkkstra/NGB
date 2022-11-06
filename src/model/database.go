package model

import (
	"byitter/src/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	schemas = []interface{}{
		&User{},
	}
)

func init() {
	connectDatabase()
	migrateSchema()
}

// connectDatabase 连接数据库
func connectDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Postgresql.Host,
		config.C.Postgresql.User,
		config.C.Postgresql.Password,
		config.C.Postgresql.Dbname,
		config.C.Postgresql.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}
}

// migrateSchema 迁移schema
func migrateSchema() {
	if err := db.AutoMigrate(schemas...); err != nil {
		panic(err)
		return
	}
}

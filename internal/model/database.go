package model

import (
	"NGB/internal/config"
	"NGB/pkg/logrus"
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

func connectDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Database.Host,
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.Dbname,
		config.C.Database.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Logger.Error(err)
		return
	}
	logrus.Logger.Info("connected to database")
}

func migrateSchema() {
	err := db.AutoMigrate(schemas...)
	if err != nil {
		logrus.Logger.Error(err)
		return
	}
	logrus.Logger.Info("migrate schema successfully")
}

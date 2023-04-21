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
		&Category{},
		&Post{},
		&UserThumbs{},
	}
)

func init() {
	connectDatabase()
	migrateSchema()
}

func connectDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Database.Sql.Host,
		config.C.Database.Sql.User,
		config.C.Database.Sql.Password,
		config.C.Database.Sql.Dbname,
		config.C.Database.Sql.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db.SetLogger(logrus.Logger)
	if err != nil {
		logrus.Logger.Error(err)
		return
	}
	// logrus.Logger.Info("PostgreSQL: connected to database")
}

func migrateSchema() {
	err := db.AutoMigrate(schemas...)
	if err != nil {
		logrus.Logger.Error(err)
		return
	}
	// logrus.Logger.Info("PostgreSQL: migrate schema successfully")
}

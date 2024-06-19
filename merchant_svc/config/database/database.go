package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"microsrvc/merchant_svc/config"
	"time"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	DB = db

	fmt.Println("[DATABASE]::CONNECTED")
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}

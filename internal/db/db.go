package db

import (
	"Soraka/conf"
	"Soraka/global"
	"Soraka/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() error {
	db, err := gorm.Open(sqlite.Open(conf.SqliteDBPath), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&model.Config{}); err != nil {
		return err
	}
	global.SqliteDB = db
	return nil
}

package model

import "gorm.io/gorm"

// Config stores simple key/value settings.
type Config struct {
	gorm.Model
	Key string `gorm:"column:k;uniqueIndex"`
	Val string `gorm:"column:v"`
}

func (Config) TableName() string {
	return "config"
}

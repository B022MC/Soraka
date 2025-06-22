package service

import (
	"Soraka/global"
	"Soraka/internal/model"
	"gorm.io/gorm/clause"
)

// SaveConfig stores or updates a config entry.
func SaveConfig(key, val string) error {
	cfg := &model.Config{Key: key, Val: val}
	return global.SqliteDB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "k"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"v": val}),
	}).Create(cfg).Error
}

// GetConfig returns the stored value for a key.
func GetConfig(key string) (string, error) {
	var cfg model.Config
	err := global.SqliteDB.Where("k = ?", key).First(&cfg).Error
	if err != nil {
		return "", err
	}
	return cfg.Val, nil
}

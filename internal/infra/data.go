package infra

import (
	conf "Soraka/internal/conf/base"
	"log"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Data 数据结构
type Data struct {
	ConfData *conf.Data
	DB       *gorm.DB
}

// NewData 初始化 Data
func NewData(confData *conf.Data, db *gorm.DB) *Data {
	return &Data{
		ConfData: confData,
		DB:       db,
	}
}

// NewDBLite 初始化 SQLite 连接实例
func NewDBLite(dataSource string, logger *log.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dataSource), &gorm.Config{
		Logger: gormLogger.New(
			zap.NewStdLog(zap.L()),
			gormLogger.Config{
				SlowThreshold: 100 * time.Millisecond,
				LogLevel:      gormLogger.Info,
				Colorful:      true,
			},
		),
		CreateBatchSize: 500,
	})
	if err != nil {
		logger.Fatalf("failed to connect to sqlite: %v", err)
	}
	logger.Printf("sqlite connect success: %s", dataSource)
	return db
}

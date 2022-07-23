package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	sysLog "log"
	"os"
	"time"
)

var db *gorm.DB

// InitDb 初始化数据库
func InitDb() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	userName := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	newLogger := logger.New(
		sysLog.New(os.Stdout, "\r\n", sysLog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond * time.Duration(10), // 慢 SQL 阈值
			LogLevel:      logger.Info,                          // LogHook level
			Colorful:      true,                                 // 彩色打印
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, userName, pass, dbName, port)
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 不要生成外键
	})
	if err != nil {
		return nil, err
	}
	db = database
	return database, nil
}

func Migrate(dbs ...interface{}) error {
	return db.AutoMigrate(dbs...)
}

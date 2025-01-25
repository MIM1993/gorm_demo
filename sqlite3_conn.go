package gorm_demo

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func InitSqlite(DBConfig *DBConfig) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBConfig.DBName), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("connect Sqlite failed: %v\n", err))
	}
	_ = db.Exec("PRAGMA journal_mode=WAL;")

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("connect sqlite failed: %v", err))
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

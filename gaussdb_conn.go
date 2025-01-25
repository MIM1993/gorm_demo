package gorm_demo

import (
	"fmt"
	gaussdb "github.com/okyer/gorm4gaussdb"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func InitGaussDB(DBConfig *DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBConfig.Host, DBConfig.Port, DBConfig.UserName, DBConfig.Password, DBConfig.DBName)
	db, err := gorm.Open(gaussdb.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("connect postgres failed: %v", err))
	}
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

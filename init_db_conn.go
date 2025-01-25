package gorm_demo

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_demo/models"
)

type DBConfig struct {
	DBType   DYTYPE
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

func NewDBConnect(dbConfig *DBConfig) (*gorm.DB, error) {
	switch dbConfig.DBType {
	case Sqlite3:
		return InitSqlite(dbConfig), nil
	case Mysql:
		return InitMySQL(dbConfig), nil
	case PostgresSql:
		return InitPostgreSQL(dbConfig), nil
	case GaussDB:
		return InitGaussDB(dbConfig), nil
	default:
		return nil, fmt.Errorf("没有这个驱动")
	}
}

func CreateTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.DemoTable{}); err != nil {
		return err
	}
	return nil
}

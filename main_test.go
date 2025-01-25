package gorm_demo

import (
	"fmt"
	"gorm_demo/models"
	"testing"
)

func TestInitSqlite(t *testing.T) {
	dbConfig := DBConfig{
		DBType:   Sqlite3,
		Host:     "127.0.0.1",
		Port:     8000,
		UserName: "root",
		Password: "123456",
		DBName:   "demo.db3",
	}
	db, err := NewDBConnect(&dbConfig)
	if err != nil {
		panic(err)
	}

	err = CreateTable(db)
	if err != nil {
		panic(err)
	}

	err = db.Model(&models.DemoTable{}).Create(&models.DemoTable{
		Name:   "穆一鸣",
		Age:    30,
		Gender: "男",
		Worker: "软件工程师",
	}).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func TestSearch(t *testing.T) {
	dbConfig := DBConfig{
		DBType:   Sqlite3,
		Host:     "127.0.0.1",
		Port:     8000,
		UserName: "root",
		Password: "123456",
		DBName:   "demo.db3",
	}
	db, err := NewDBConnect(&dbConfig)
	if err != nil {
		panic(err)
	}
	data := models.DemoTable{}
	err = db.Model(&models.DemoTable{}).Find(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", data)
}

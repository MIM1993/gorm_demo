package gorm_demo

type DYTYPE int8

const (
	Sqlite3 DYTYPE = iota + 1
	Mysql
	PostgresSql
	GaussDB
)

package models

type DemoTable struct {
	ID     int64  `json:"id"                  gorm:"primary_key;auto_increment;comment:主键"`
	Name   string `json:"node_name"           gorm:"default:''"`
	Age    int    `json:"age"                 gorm:"default:0"`
	Gender string `json:"gender"              gorm:"default:''"`
	Worker string `json:"worker"              gorm:"default:''"`
}

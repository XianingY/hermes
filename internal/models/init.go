package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:yang0725@tcp(127.0.0.1:3306)/hermes_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&RoomBasic{}, &UserBasic{}, &UserOfRoom{})
	DB = db

}

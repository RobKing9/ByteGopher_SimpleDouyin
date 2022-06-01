package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
/*
	dao包 用来封装数据库相关的操作
*/

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}


func InitMysql() {
	dsn := "root:Byt3G0pheR51522zzwlwlbb@tcp(121.40.120.222:43306)/simpledouyin?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb ,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = gormdb
}


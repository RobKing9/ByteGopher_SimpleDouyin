package dao

/*
	dao包
用来封装数据库相关的操作
*/

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlDb *gorm.DB

func InitMysql(){
	var err error
	dsn := "root:Byt3G0pheR51522zzwlwlbb@tcp(121.40.120.222:43306)/simpledouyin?charset=utf8mb4&parseTime=True&loc=Local"
	// MysqlDb, err = gorm.Open("mysql", dsn)
	db, err := gorm.Open(mysql.Open(dsn),  &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("gorm open failed:", err.Error())
	}
	fmt.Println("gorm open success")
	// fmt.Println(MysqlDb)
	MysqlDb = db 
}

func GetDB() *gorm.DB {
	return MysqlDb
}

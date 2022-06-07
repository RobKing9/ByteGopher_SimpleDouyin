package dao

/*
	dao包
用来封装数据库相关的操作
*/

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDb *gorm.DB

func InitMysql() *gorm.DB {
	dsn := "root:Byt3G0pheR51522zzwlwlbb@tcp(121.40.120.222:43306)/simpledouyin?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDb, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("gorm open failed:", err.Error())
	}
	fmt.Println("gorm open success")
	// fmt.Println(MysqlDb)
	return MysqlDb
}

func GetDB() *gorm.DB {
	return MysqlDb
}

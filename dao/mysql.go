package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Db *gorm.DB

func InitMysql() *gorm.DB {
	dsn := "root:Byt3G0pheR51522zzwlwlbb@tcp(121.40.120.222:43306)/simpledouyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("gorm open failed:", err.Error())
		panic(err)
	}
	return db
}

func setup() {
	Db = InitMysql()
	Db.DB().SetMaxIdleConns(10)                   //最大空闲连接数
	Db.DB().SetMaxOpenConns(30)                   //最大连接数
	Db.DB().SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	Db.LogMode(true)
}

func GetDB() *gorm.DB {
	if err := Db.DB().Ping(); err != nil {
		Db.Close()
		//Db = InitMysql()
		setup()
	}
	return Db
}

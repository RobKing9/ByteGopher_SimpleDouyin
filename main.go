package main

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/router"
	"github.com/gin-gonic/gin"
)

/*
	项目入口
*/

func main() {
	//初始化数据库
	dao.InitMysql()
	//设置路由
	r := gin.Default()
	r = router.CollectRouter(r)
	//启动路由
	r.Run(":")
}

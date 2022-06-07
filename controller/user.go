package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Info 用户信息
func Info(c *gin.Context) {
	fmt.Println("Info")
	// users, _  := model.GetAllUserModels()
	// fmt.Println("GetAllUserModels") 
	// c.JSON(http.StatusOK, users)
}

// Login 用户登录
func Login(c *gin.Context) {

}

// Register 用户注册
func Register(c *gin.Context) {

}

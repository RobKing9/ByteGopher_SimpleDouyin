package middleware

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
	middleware包 存放中间件
*/

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		fmt.Print("请求token", tokenString)

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		DB := dao.GetDB()
		var u model.User
		DB.First(&u, userId)

		// 验证用户是否存在
		if u.UserId == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		c.Set("user", u)

		c.Next()
	}
}

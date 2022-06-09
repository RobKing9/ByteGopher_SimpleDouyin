package middleware

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"github.com/gin-gonic/gin"
	"log"
)

/*
	middleware包 存放中间件
*/

type JwtResponse struct {
	StatusCode int64  `json:status_code`
	StatusMsg  string `json:status_msg`
}

func AuthMiddleware() gin.HandlerFunc {
	//fmt.Println("进入中间件了")
	return func(c *gin.Context) {
		var authFlag = false

		tokenString := c.Query("token")
		//log.Printf("请求token:%s", tokenString)

		//validate token formate
		if tokenString == "" {
			//c.JSON(http.StatusUnauthorized, JwtResponse{401, "token为空"})
			log.Println("token为空")
			//c.Abort()
			c.Set("flag", authFlag)
			return
		}

		//tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)

		if err != nil || !token.Valid {
			//c.JSON(http.StatusUnauthorized, JwtResponse{401, "权限不足"})
			log.Println("权限不足:", err.Error())
			//c.Abort()
			c.Set("flag", authFlag)
			return
		}

		//token通过验证, 获取claims中的UserID
		userId := claims.UserId
		var u model.UserModel
		dao.MysqlDb.Where("user_id=?", userId).First(&u)

		// 验证用户是否存在
		if u.UserID == 0 {
			//c.JSON(http.StatusUnauthorized, JwtResponse{401, "权限不足"})
			log.Println("用户不存在")
			//c.Abort()
			c.Set("flag", authFlag)
			return
		}

		// flag = true
		authFlag = true
		c.Set("flag", authFlag)

		//用户存在 将user信息写入上下文
		c.Set("user", u)
		c.Next()
	}
}

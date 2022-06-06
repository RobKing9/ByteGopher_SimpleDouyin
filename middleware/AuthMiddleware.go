package middleware

/*
	middleware包 存放中间价
*/

import (
	"net/http"
	"strconv"
	"strings"

	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils/jwtTool"
)

const (
	contextKeyUserObj = "authedUserObj"
	bearerLength      = len("Bearer ")
)

// JwtAuthWithUserId 从 url-query 的 token 获取 JWTString 或者从请求头 Authorization 中获取 JWTString
//
// jwt_v1.JwtParseUser(token)解析 JWTString 可以获取 model.User 结构体
//
// 设置用户信息到 gin.Context 其他的 handler 通过 gin.Context.Get(contextKeyUserObj),
// 在进行用户 Type Assert 得到 model.User 结构体
//
// 使用了 jwt-middle 之后的 handle 从 gin.Context 中获取用户信息
func JwtAuthWithUserId() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 从 url 中获取 user_id
		userId, err := strconv.ParseInt(c.Query("user_id"), 0, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed,
				model.Response{
					StatusCode: model.SCodeFalse,
					StatusMsg:  "no userid",
				})
			return
		}

		token, ok := c.GetQuery("token")
		if !ok { // 或者从请求头 Authorization 中获取 JWTString
			hToken := c.GetHeader("Authorization")
			if len(hToken) < bearerLength {
				c.AbortWithStatusJSON(http.StatusPreconditionFailed,
					model.Response{
						StatusCode: model.SCodeFalse,
						StatusMsg:  "header Authorization has not Bearer token",
					})
				return
			}
			token = strings.TrimSpace(hToken[bearerLength:])
		}

		// 解析 token ，返回 user 结构体
		usr, err := jwtTool.JwtParseUser(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed,
				model.Response{
					StatusCode: model.SCodeFalse,
					StatusMsg:  "no permission",
				})
			return
		}

		// 认证用户: userid 与 token 是否一致
		if (usr.UserId & userId) != userId {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed,
				model.Response{
					StatusCode: model.SCodeFalse,
					StatusMsg:  "userid not match token",
				})
			return
		}

		// 将用户模型存储在 gin 上下文中
		c.Set(contextKeyUserObj, *usr)
		c.Next()
	}
}

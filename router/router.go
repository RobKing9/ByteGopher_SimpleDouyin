package router

import (
	"ByteGopher_SimpleDouyin/controller"
	"ByteGopher_SimpleDouyin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	router包 用来封装路由
*/
func CollectRouter(r *gin.Engine) *gin.Engine {
	douyin := r.Group("/douyin")
	{
		/*test*/
		douyin.GET("/hello", func(c *gin.Context) {
			//将发送的信息封装成JSON发送给浏览器
			c.JSON(http.StatusOK, gin.H{
				//这是我们定义的数据
				"message": "hi",
			})
		})
		/*
			基础接口
		*/
		//视频流接口
		douyin.GET("/feed", controller.NewVideoController().Feed)
		//用户组
		user := douyin.Group("/user")
		{
			userController := controller.NewUserController()
			//用户信息
			user.GET("/", middleware.AuthMiddleware(), userController.Info)
			//用户登录接口
			user.POST("/login/", userController.Login)
			//用户注册接口
			user.POST("/register/", userController.Register)
		}
		//发布视频接口
		publish := douyin.Group("/publish", middleware.AuthMiddleware())
		{
			publish.GET("/action")
			publish.GET("/list")
		}

		/*
			扩展接口-I
		*/
		//点赞接口
		favorite := douyin.Group("/favorite")
		{
			//赞操作
			favorite.GET("/action")
			//赞列表
			favorite.GET("/list")
		}
		//评论接口
		comment := douyin.Group("/comment")
		{
			//评论操作
			comment.GET("/action")
			//评论列表
			comment.GET("/list")
		}

		/*
			扩展接口—II
		*/
		relation := douyin.Group("/relation")
		{
			relation.GET("/action")
			relation.GET("/follow/list")
			relation.GET("/follower/list")
		}
	}
	return r
}

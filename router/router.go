package router

import (
	"ByteGopher_SimpleDouyin/controler"
	"ByteGopher_SimpleDouyin/middleware"
	"github.com/gin-gonic/gin"
)

/*
	router包 用来封装路由
*/
func CollectRouter(r *gin.Engine) *gin.Engine {
	douyin := r.Group("/douyin")
	{
		/*
			基础接口
		*/
		//视频流接口
		douyin.GET("/feed")
		//用户组
		user := douyin.Group("/user")
		{
			//用户信息
			user.GET("/", controler.Info)
			//用户登录接口
			user.POST("/login", controler.Login)
			//用户注册接口
			user.POST("/register", controler.Register)
		}
		//发布视频接口
		publish := douyin.Group("/publish")
		{
			publish.GET("/action")
			publish.GET("/list",middleware.JwtAuthWithUserId(),controler.PublishList)
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
		relation.Use(middleware.JwtAuthWithUserId())
		{
			relation.POST("/action",controler.Action)
			relation.GET("/follow/list",controler.FollowList)
			relation.GET("/follower/list",controler.FollowerList)
		}
	}
	return r
}

package controler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
)

type RelationActionResponse struct {
	model.Response
}

// Action 处理关注操作
func Action(ctx *gin.Context) {
	userId,err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
		return
	}

	toUserId,err := strconv.ParseInt(ctx.Query("to_user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no this user",
			},
		})
		return
	}

	actionType,err := strconv.ParseInt(ctx.Query("action_type"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "illegal operation",
			},
		})
		return
	}


	// 执行操作
	// 查询该用户是否存在/ 缓存？数据库

	// 关注 or 取消关注
	// 判断 action_type  = 1 关注， = 2 取消关注
	// 操作 redis 缓存数据（redis 是否需要持久化）
	err = dao.Follow(toUserId, userId, int32(actionType))
	if err != nil {
		ctx.JSON(http.StatusOK,RelationActionResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "illegal operation",
			},
		})
	}

	ctx.JSON(http.StatusOK,RelationActionResponse{
		Response:model.Response{
			StatusCode: model.SCodeSuccess,
		},
	})
}





type RelationFansListResponse struct {
	model.Response
	Users	[]model.User	`json:"user_list"`
}


// FollowerList 获取粉丝列表
func FollowerList(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFansListResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
	}

	// 查询粉丝列表，返回粉丝列表（注意加锁）
	users := dao.GetFansList(dao.UserIdType(userId))


	ctx.JSON(http.StatusOK,RelationFansListResponse{
		Response:model.Response{
			StatusCode: model.SCodeSuccess,
		},

		Users:users,
	})
}





type RelationFollowListResponse struct {
	model.Response
	Users	[]model.User	`json:"user_list"`
}

// FollowList 获取关注列表
func FollowList(ctx *gin.Context) {

	userId, err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFollowListResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
	}


	// 查询关注列表，返回关注列表（注意加锁）
	users := dao.GetFollowList(dao.UserIdType(userId))


	ctx.JSON(http.StatusOK,RelationFollowListResponse{
		Response:model.Response{
			StatusCode: model.SCodeSuccess,
		},

		Users:users,
	})
}
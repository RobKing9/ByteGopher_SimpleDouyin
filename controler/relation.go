package controler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/dao/cache"
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


	// 关注 or 取消关注
	// 判断 action_type  = 1 关注， = 2 取消关注
	// 操作 redis 缓存数据（redis 是否需要持久化）
	err = followAction(userId, toUserId, int32(actionType))
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


// 关注操作
func followAction(myId, toUserId int64, actionType int32) error {

	// 自身用户信息是否存在
	myInfo := model.User{}
	// 优先从缓存中获取自身信息
	myInfo = cache.JsonToStruct(myId)
	if myInfo.UserId == -1 {	// 缓存中无该用户自身信息
		// 再从数据库中获取
		_,err := dao.NewUserModel().SearchUserById(myId)
		if err != nil {
			return err
		}
	}


	switch actionType {
	case 1:		// 关注
		err := cache.FollowAction(toUserId, myId)
		if err != nil {
			return err
		}
	case 2:		// 取消关注
		err := cache.CancelFollowAction(toUserId,myId)
		if err != nil {
			return err
		}
	default:
		break
	}

	return nil
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
	users := cache.GetFansListByUserId(userId)


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
	users := cache.GetFollowListByUserId(userId)


	ctx.JSON(http.StatusOK,RelationFollowListResponse{
		Response:model.Response{
			StatusCode: model.SCodeSuccess,
		},

		Users:users,
	})
}
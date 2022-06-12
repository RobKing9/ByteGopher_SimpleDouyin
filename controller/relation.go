package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ByteGopher_SimpleDouyin/model"
)

// RelationActionResponse 关注操作返回结构体
type RelationActionResponse struct {
	model.RespModel
}

// RelationAction 处理关注操作
func RelationAction(ctx *gin.Context) {

	// 若鉴权失败，则返回
	flag,exist := ctx.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: this flag not exist")
	}
	if !flag.(bool) {	// 认证失败
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
		return
	}

	// 获取自身 userId
	userId,exist := ctx.Get("userid")
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: user not exist")
	}
	//userId,err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	//if err != nil {
	//	ctx.JSON(http.StatusOK,	RelationActionResponse{
	//		RespModel:model.RespModel{
	//			StatusCode: model.SCodeFalse,
	//			StatusMsg: err.Error(),
	//		},
	//	})
	//	return
	//}

	// 获取对方 toUserId
	toUserId,err := strconv.ParseInt(ctx.Query("to_user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: err.Error(),
			},
		})
		return
	}

	// 获取关注操作类型 action_type
	actionType,err := strconv.ParseInt(ctx.Query("action_type"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: err.Error(),
			},
		})
		return
	}


	// 关注 or 取消关注
	// 判断 action_type  = 1 关注， = 2 取消关注
	err = followAction(userId.(int64), toUserId, int32(actionType))
	if err != nil {
		ctx.JSON(http.StatusOK,RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK,RelationActionResponse{
		RespModel:model.RespModel{
			StatusCode: model.SCodeSuccess,
		},
	})
}


// 关注操作
// 判断 action_type  = 1 关注， = 2 取消关注
func followAction(myId, toUserId int64, actionType int32) error {

	switch actionType {
	case 1:		// 关注
		//err := cache.FollowAction(toUserId, myId)
		err := dao.SaveFollowInToTable(myId,toUserId)
		if err != nil {
			return err
		}
	case 2:		// 取消关注
		//err := cache.CancelFollowAction(toUserId,myId)
		err := dao.DeleteFansInToTable(myId,toUserId)
		if err != nil {
			return err
		}
	default:
		return errors.New("illegal operation")
	}

	return nil
}




// RelationFansListResponse 返回粉丝列表结构体
type RelationFansListResponse struct {
	model.RespModel
	Users	[]model.User	`json:"user_list"`
}


// FollowerList 获取粉丝列表
func FollowerList(ctx *gin.Context) {
	// 若鉴权失败，则返回
	flag,exist := ctx.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("FollowerList: this flag not exist")
	}
	if !flag.(bool) {	// 认证失败
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
		return
	}

	myId, err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFansListResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
	}

	// 查询粉丝列表，返回粉丝列表的切片（注意redis加锁）
	//users := cache.GetFansListByUserId(myId)
	users, err := dao.GetFansList(myId)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFollowListResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: err.Error(),
			},
		})
	}

	ctx.JSON(http.StatusOK,RelationFansListResponse{
		RespModel:model.RespModel{
			StatusCode: model.SCodeSuccess,
		},

		Users:users,
	})
}




// RelationFollowListResponse 返回关注列表的结构体
type RelationFollowListResponse struct {
	model.RespModel
	Users	[]model.User	`json:"user_list"`
}

// FollowList 获取关注列表
func FollowList(ctx *gin.Context) {
	// 若鉴权失败，则返回
	flag,exist := ctx.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("FollowList: this flag not exist")
	}
	if !flag.(bool) {	// 认证失败
		ctx.JSON(http.StatusOK,	RelationActionResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})

		return
	}

	myId, err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFollowListResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: "no permission",
			},
		})
	}


	// 查询关注列表，返回关注列表切片（注意redis加锁）
	//users := cache.GetFollowListByUserId(myId)
	users,err := dao.GetFollowList(myId)
	if err != nil {
		ctx.JSON(http.StatusOK, RelationFollowListResponse{
			RespModel:model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg: err.Error(),
			},
		})
	}


	ctx.JSON(http.StatusOK,RelationFollowListResponse{
		RespModel:model.RespModel{
			StatusCode: model.SCodeSuccess,
		},

		Users:users,
	})
}
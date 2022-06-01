package controler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
)


//type PublishListRequest struct {
//	UserId	int64	`json:"user_id"`
//	Token	string	`json:"token"`
//}

type PublishListResponse struct {
	model.Response                      // 返回状态
	VideoList	[]dao.VideoModel `json:"video_list"`
}


func PublishList(ctx *gin.Context) {

	// 自定义中间件 middleware。JwtAuthWithUserId 会检验token合法性
	// 并且会认证用户身份

	//// 获取 Query 参数
	//token := ctx.Query("token")
	usrId, err := strconv.ParseInt(ctx.Query("user_id"),0,64)
	if err != nil{
		ctx.JSON(http.StatusOK,PublishListResponse{
			Response:model.Response  {
				StatusCode: model.SCodeFalse,
			},
		})
		return
	}
	//
	//// 解析验证 token 是否合法
	//user, err := jwtTool.JwtParseUser(token)
	//if err != nil {
	//	ctx.JSON(http.StatusOK,PublishListResponse{
	//		Response:Response{
	//			StatusCode: SCodeFalse,
	//			StatusMsg: err.Error(),
	//		},
	//	})
	//	return
	//}
	//
	//// 身份认证，是否一致
	//if (user.UserId & usrId) != usrId {
	//	ctx.JSON(http.StatusOK,PublishListResponse{
	//		Response:Response{
	//			StatusCode: SCodeFalse,
	//			StatusMsg: "token does not match userId",
	//		},
	//	})
	//	return
	//}


	// 获取视频列表数据
	videos,err := dao.NewVideoModel().Search(usrId)
	if err != nil{
		ctx.JSON(http.StatusOK,PublishListResponse{
			Response:model.Response  {
				StatusCode: model.SCodeFalse,
			},
		})
		return
	}


	// 返回响应
	ctx.JSON(http.StatusOK, PublishListResponse{
		Response:model.Response  {
			StatusCode: model.SCodeSuccess,
		},
		VideoList: videos,
	})
}
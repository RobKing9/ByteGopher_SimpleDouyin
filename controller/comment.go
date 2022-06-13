package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type CommentController interface {
	CommentAction(c *gin.Context)
	CommentList(c *gin.Context)
}
type commentController struct {
	comDao  dao.CommentDao
	userDap dao.UserDao
}

func NewCommentController() CommentController {
	return &commentController{comDao: dao.NewCommentDao()}
}

// CommentListResponse
// 评论列表返回参数
type CommentListResponse struct {
	StatusCode  int32               `json:"status_code"`
	StatusMsg   string              `json:"status_msg,omitempty"`
	CommentList []model.CommentInfo `json:"comment_list,omitempty"`
}

// CommentActionResponse
// 发表评论返回参数
type CommentActionResponse struct {
	StatusCode int32             `json:"status_code"`
	StatusMsg  string            `json:"status_msg,omitempty"`
	Comment    model.CommentInfo `json:"comment"`
}

// CommentAction
// 发表 or 删除评论 comment/action/
func (com *commentController) CommentAction(c *gin.Context) {
	log.Println("CommentController-Comment_Action: running") //函数已运行
	// 若鉴权失败，则返回
	flag, exist := c.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: this flag not exist")
	}
	if !flag.(bool) { // 认证失败
		c.JSON(http.StatusOK, RelationActionResponse{
			RespModel: model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg:  "no permission",
			},
		})
		return
	}
	//获取userId
	userid, exit := c.Get("userid")
	if exit != true {
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment userId json invalid",
		})
		log.Println("CommentController-Comment_Action: return comment userId json invalid") //函数返回userId无效
		return
	}
	//获取videoId
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	//错误处理
	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment videoId json invalid",
		})
		log.Println("CommentController-Comment_Action: return comment videoId json invalid") //函数返回视频id无效
		return
	}
	//获取操作类型
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	//错误处理
	if err != nil || actionType < 1 || actionType > 2 {
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: -1,
			StatusMsg:  "comment actionType json invalid",
		})
		log.Println("CommentController-Comment_Action: return actionType json invalid") //评论类型数据无效
		return
	}
	if actionType == 1 { //actionType为1，则进行发表评论操作
		content := c.Query("comment_text")
		//发表评论数据准备
		var sendComment model.CommentModel
		rand.Seed(time.Now().UnixNano())
		id := rand.Int63() // 生成比较大的随机数
		sendComment.CommentID = id
		sendComment.UserID = userid.(int64)
		sendComment.VideoID = videoId
		sendComment.Content = content
		timeNow := time.Now()
		sendComment.CreateDate = timeNow
		//发表评论
		res, err := com.comDao.InsertComment(sendComment)
		fmt.Println("插入评论:", res)
		//User, err := com.userDap.GetCommonUserByID(userid.(int64))
		fmt.Println("User:", userid.(int64))
		commentInfo := model.CommentInfo{
			UserInfo: model.User{
				Id: userid.(int64),
			},
			Content:    content,
			CreateDate: timeNow.String(),
		}
		//发表评论失败
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "send comment failed",
			})
			log.Println("CommentController-Comment_Action: return send comment failed") //发表失败
			return
		}

		//发表评论成功:
		//返回结果
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "send comment success",
			Comment:    commentInfo,
		})
		log.Println("CommentController-Comment_Action: return Send success") //发表评论成功，返回正确信息
		return
	} else { //actionType为2，则进行删除评论操作
		//获取要删除的评论的id
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  "delete commentId invalid",
			})
			log.Println("CommentController-Comment_Action: return commentId invalid") //评论id格式错误
			return
		}
		//删除评论操作
		err = com.comDao.DeleteComment(commentId)
		if err != nil { //删除评论失败
			str := err.Error()
			c.JSON(http.StatusOK, CommentActionResponse{
				StatusCode: -1,
				StatusMsg:  str,
			})
			log.Println("CommentController-Comment_Action: return delete comment failed") //删除失败
			return
		}
		//删除评论成功
		c.JSON(http.StatusOK, CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "delete comment success",
		})

		log.Println("CommentController-Comment_Action: return delete success") //函数执行成功，返回正确信息
		return
	}
}

// CommentList
// 查看评论列表 comment/list/
func (com *commentController) CommentList(c *gin.Context) {
	log.Println("CommentController-Comment_List: running") //函数已运行
	// 若鉴权失败，则返回
	flag, exist := c.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: this flag not exist")
	}
	if !flag.(bool) { // 认证失败
		c.JSON(http.StatusOK, RelationActionResponse{
			RespModel: model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg:  "no permission",
			},
		})
		return
	}
	////获取userId
	//id, _ := c.Get("userId")
	//userid, _ := id.(string)
	//userId, err := strconv.ParseInt(userid, 10, 64)
	////log.Printf("err:%v", err)
	////log.Printf("userId:%v", userId)

	//获取videoId
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	//错误处理
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  "comment videoId json invalid",
		})
		log.Println("CommentController-Comment_List: return videoId json invalid") //视频id格式有误
		return
	}
	log.Printf("videoId:%v", videoId)

	//评论列表
	commentListModel, err := com.comDao.GetCommentList(videoId)
	commentList := []model.CommentInfo{}
	for _, v := range commentListModel {
		//User, _ := com.userDap.GetCommonUserByID(v.UserID)
		commentInfo := model.CommentInfo{
			Id: v.CommentID,
			UserInfo: model.User{
				Id: v.UserID,
			},
			Content:    v.Content,
			CreateDate: v.CreateDate.String(),
		}
		commentList = append(commentList, commentInfo)
	}
	if err != nil { //获取评论列表失败
		c.JSON(http.StatusOK, CommentListResponse{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		})
		log.Println("CommentController-Comment_List: return list false") //查询列表失败
		return
	}

	//获取评论列表成功
	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "get comment list success",
		CommentList: commentList,
	})
	log.Println("CommentController-Comment_List: return success") //成功返回列表
	return
}

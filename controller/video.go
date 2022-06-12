package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Feed(c *gin.Context)
	PublishAction(c *gin.Context)
	PublishList(c *gin.Context)
}

type videoController struct {
	videoDao dao.VideoDao
}

func NewVideoController() VideoController {
	return &videoController{
		videoDao: dao.NewVideoDao(),
	}
}

func (controller *videoController) Feed(c *gin.Context) {
	// latest_time, ok := c.GetQuery("latest_time")
	// token, ok := c.GetQuery("token")
	videoList := make([]model.Video, 0)
	// var video model.Video
	// var user model.User
	videos, err := controller.videoDao.GetVideos()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	for _, v := range videos {

		user := model.User{
			Id:            v.Author.UserID,
			Name:          v.Author.UserName,
			FollowCount:   v.Author.FollowCount,
			FollowerCount: v.Author.FollowerCount,
			IsFollow:      false,
		}
		video := model.Video{
			Id:            v.VideoID,
			Author:        user,
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
		}
		fmt.Println(video)
		videoList = append(videoList, video)
	}
	feed := model.FeedResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取视频成功",
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}

	fmt.Println(feed.VideoList)
	c.JSON(http.StatusOK, feed)
}

func makeVideoModel(id int64, user_id int64, playUrl string, coverURl string, favoriteCount int64, commentCount int64, title string) model.VideoModel {
	return model.VideoModel{
		VideoID:       id,
		UserID:        user_id,
		PlayURL:       playUrl,
		CoverURL:      coverURl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		Title:         title,
		PublishTime:   time.Now(),
	}
}

func (controller *videoController) PublishAction(c *gin.Context) {
	// 认证失败
	flag, _ := c.Get("flag")
	if !flag.(bool) {
		c.JSON(http.StatusUnauthorized, model.RespModel{
			StatusCode: -1,
			StatusMsg:  "请先登录!!!",
		})
		log.Println("请先登录！")
		return
	}

	data, _ := c.FormFile("data")
	title := c.PostForm("title")
	//tokenString := c.PostForm("token")
	id := utils.RandRangeIn(10000000, 99999999) // 随机数生成视频ID
	var FavoriteCount int64 = 0
	var CommentCount int64 = 0
	var testVideo *model.VideoModel

	// TODO: 判断上传的视频文件是否为mp4格式

	// 获取user
	//user, isExist := c.Get("user")
	//if isExist != true {
	//	res := model.RespModel{
	//		StatusCode: -1,
	//		StatusMsg:  "用户不存在",
	//	}
	//	c.JSON(200, res)
	//	return
	//}

	// 自己生成user
	//_, claims, err := utils.ParseToken(tokenString)
	//userId := claims.UserId
	//var u model.UserModel
	//dao.MysqlDb.Where("user_id=?", userId).First(&u)

	// 取到userid
	userId, _ := c.Get("userid")

	// 判断id是否存在 若已存在则重新生成
	for {
		testVideo, _ = controller.videoDao.GetVideoModelByID(id)
		if testVideo == nil { // nil说明record not found 也就是说id不存在 不需要重新生成
			break
		}
		id = utils.RandRangeIn(10000000, 99999999)
	}
	// 视频先保存到本地
	c.SaveUploadedFile(data, "/tmp/${id}.mp4")

	//TODO: 获取封面url 目前封面暂定一个固定的图
	coverUrl := "http://cdn1.pic.y1ng.vip/uPic/IMG_3567.JPG"

	// 上传至七牛云
	retKey, err := utils.QiniuUpload("video/${id}.mp4", "/tmp/${id}.mp4")

	// 无论远程oss上传是否成功 都删除本地文件
	utils.DeleteFile("/tmp/${id}.mp4")

	// 如果上传失败则返回
	if err != nil {
		res := model.RespModel{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(200, res)
		return
	}

	video := makeVideoModel(int64(id), userId.(int64), retKey, coverUrl, FavoriteCount, CommentCount, title)

	// 插入数据到数据库
	err = controller.videoDao.AddVideoModel(&video)
	if err != nil {
		res := model.RespModel{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(200, res)
		return
	}

	// 返回成功
	c.JSON(200, model.RespModel{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

// 获取用户发布的视频列表
func (controller *videoController) PublishList(c *gin.Context) {
	// 认证失败
	flag, _ := c.Get("flag")
	if !flag.(bool) {
		c.JSON(http.StatusUnauthorized, model.RespVideoList{
			StatusCode: -1,
			StatusMsg:  "请先登录!!!",
			VideoList:  nil,
		})
		log.Println("请先登录！")
		return
	}

	userID := c.PostForm("user_id")
	userIDint, _ := strconv.ParseInt(userID, 10, 64)
	videoModelList, _ := controller.videoDao.GetVideoListByUserId(userIDint)
	videoList := utils.ConvertVideoModelListToVideoList(videoModelList)

	res := model.RespVideoList{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
	}
	c.JSON(200, res)
}



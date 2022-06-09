package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Feed(c *gin.Context)

}

type videoController struct {
	videoDao dao.VideoDao
}

func NewVideoController() VideoController {
	return &videoController{
		videoDao: dao.NewVideoDao(),
	}
}

func (controller *videoController)Feed(c *gin.Context) {
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
				Id: v.Author.UserID,
				Name: v.Author.UserName,
				FollowCount: v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				IsFollow: false,
			}
			video := model.Video{
				Id: v.VideoID,
				Author: user,
				PlayUrl: v.PlayURL,
				CoverUrl: v.CoverURL,
				FavoriteCount: v.FavoriteCount,
				CommentCount: v.CommentCount,
				IsFavorite: true,
			}
			fmt.Println(video)
			videoList = append(videoList, video)
	}
	feed := model.FeedResponse{
		Response:  model.Response{
			StatusCode: 0,
			StatusMsg: "获取视频成功",
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}

	fmt.Println(feed.VideoList)
	c.JSON(http.StatusOK, feed)
}

package service

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"fmt"
	"time"
)

type VideoService interface {
		GetFeed() (model.FeedResponse, error)
}

type videoService struct {
		userDao dao.UserDao
		videoDao dao.VideoDao
}

func NewVideoService() VideoService {
		return &videoService{
			userDao: dao.NewUserDao(),
			videoDao: dao.NewVideoDao(),
		}
}

func (service videoService)GetFeed() (model.FeedResponse, error) {
	
	videoList := make([]model.Video, 0)
	// var video model.Video
	// var user model.User
	videos, _ := service.videoDao.GetVideoModels()
	for _, v := range videos {
			u, _ := service.userDao.GetUserModelByID( int(v.UserID.Int64))
			user := model.User{
				Id: int64(u.UserID),
				Name: u.UserName.String,
				FollowCount: u.FollowCount.Int64,
				FollowerCount: u.FollowerCount.Int64,
				IsFollow: false,
			}
			video := model.Video{
				Id: int64(v.VideoID),
				Author: user,
				PlayUrl: v.PlayURL.String,
				CoverUrl: v.CoverURL.String,
				FavoriteCount: v.FavoriteCount.Int64,
				CommentCount: v.CommentCount.Int64,
				IsFavorite: true,
			}
			fmt.Println(video)
			videoList = append(videoList, video)
	}
	return model.FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}, nil
}
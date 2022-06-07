package service

import (
	"ByteGopher_SimpleDouyin/model"
	"time"

	"github.com/jinzhu/gorm"
)

// type feedService struct {
// 	videoModel model.VideoModel
// 	userModel  model.UserModel
// 	// followModel model.FollowModel
// }

// func NewFeedService(videoModel model.VideoModel, userModel model.UserModel) *feedService {
// 	return &feedService{
// 		videoModel: videoModel,
// 		userModel: userModel,
// 	}
// }

type VideoService interface {
		GetFeed() (model.FeedResponse, error)
}

type videoService struct {
		db *gorm.DB
}

func NewVideoService(db *gorm.DB) VideoService {
		return &videoService{db: db}
}

func (service videoService)GetFeed() (model.FeedResponse, error) {
	
	videoList := make([]model.Video, 0)
	// var video model.Video
	// var user model.User
	videos, _ := model.GetVideoModels(service.db)
	for _, v := range videos {
			u, _ := model.GetUserModelByID(service.db, int(v.UserID.Int64))
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
			videoList = append(videoList, video)
	}
	return model.FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}, nil
}
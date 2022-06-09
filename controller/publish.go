package controller

import (
	"ByteGopher_SimpleDouyin/model"
)

// upload video的返回
type RespModel struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 视频列表的返回
type RespVideoList struct {
	StatusCode int64         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	VideoList  []model.Video `json:"video_list"` // 用户发布的视频列表
}

// 根据视频信息，生成视频的结构体对象
func makeVideo(id int64, author model.User, playUrl string, coverURl string, favoriteCount int64, commentCount int64, isFavorite bool, title string) model.Video {
	return model.Video{
		Id:            id,
		Author:        author,
		PlayUrl:       playUrl,
		CoverUrl:      coverURl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         title,
	}
}

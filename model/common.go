package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

// omitempty会舍弃false
type User struct {
	Id            int64  `gorm:"column:user_id;primaryKey;unique;not null;autoIncrement" json:"id,omitempty"`
	Name          string `gorm:"column:user_name" json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// func (Video) TableName() string {
// 	return "video"
// }

// func (User) TableName() string {
// 	return "user"
// }

// upload video的返回
type RespModel struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 视频列表的返回
type RespVideoList struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"` // 用户发布的视频列表
}

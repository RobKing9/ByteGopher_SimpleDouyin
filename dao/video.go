package dao

import (
	"fmt"

	"ByteGopher_SimpleDouyin/model"
)

type VideoModel struct {
	Video 		model.Video		 `gorm:"embedded"`
	//UserId		int64		`json:"-" gorm:"-"`
	//PublishTime	time.Time		`gorm:"publish_time"`
}


func (v VideoModel) TableName() string {
	return "video"
}

func NewVideoModel() *VideoModel {
	return &VideoModel{}
}

func (v *VideoModel) Save(obj interface{}) error {

	return nil
}

func (v *VideoModel) Search(usrId int64) ([]VideoModel,error) {
	video := []VideoModel{}

	str := fmt.Sprintf("video.video_id,u.*" +
		",video.play_url,video.cover_url,video.favorite_count,video.comment_count,video.title")

	tx := GetDB().Select(str).Joins("join user u using(user_id)").Table("video").Find(&video)
	return video,tx.Error
}
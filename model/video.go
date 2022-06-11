package model

import (
	"time"
)

type VideoModel struct {
	VideoID       int64     `gorm:"column:video_id;primaryKey;unique;not null;autoIncrement" json:"video_id"`
	UserID        int64     `gorm:"column:user_id" json:"user_id"`
	Author        UserModel `json:"author" gorm:"foreignKey:UserID"`
	PlayURL       string    `gorm:"column:play_url" json:"play_url"`
	Title         string    `gorm:"column:title" json:"title"`
	CoverURL      string    `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count" json:"comment_count"`
	PublishTime   time.Time `gorm:"column:publish_time" json:"publish_time"`
}

// TableName sets the insert table name for this struct type
func (VideoModel) TableName() string {
	return "video"
}

// func AddVideoModel(m *VideoModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func DeleteVideoModelByID(id int) (bool, error) {
// 	if err := dao.MysqlDb.Delete(&VideoModel{}, id).Error; err != nil {
// 		return false, err
// 	}
// 	return dao.MysqlDb.RowsAffected > 0, nil
// }

// func DeleteVideoModel(condition string, args ...interface{}) (int64, error) {
// 	if err := dao.MysqlDb.Where(condition, args...).Delete(&VideoModel{}).Error; err != nil {
// 		return 0, err
// 	}
// 	return dao.MysqlDb.RowsAffected, nil
// }

// func UpdateVideoModel(m *VideoModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func GetVideoModelByID(id int) (*VideoModel, error) {
// 	var m VideoModel
// 	if err := dao.MysqlDb.First(&m, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &m, nil
// }

// func GetVideoModels() ([]*VideoModel, error) {
// 	res := make([]*VideoModel, 0)
// 	fmt.Println(dao.MysqlDb)
// 	if err := dao.MysqlDb.Table("video").Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(" GetVideoModels success")
// 	return res, nil
// }

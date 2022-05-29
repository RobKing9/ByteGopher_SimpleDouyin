package model

import (
	"database/sql"
	"time"
)

type VideoModel struct {
	VideoID       int            `gorm:"column:video_id;primaryKey;unique;not null;autoIncrement" json:"video_id"`
	UserID        sql.NullInt64  `gorm:"column:user_id" json:"user_id"`
	PlayURL       sql.NullString `gorm:"column:play_url" json:"play_url"`
	Title         sql.NullString `gorm:"column:title" json:"title"`
	CoverURL      sql.NullString `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount sql.NullInt64  `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  sql.NullInt64  `gorm:"column:comment_count" json:"comment_count"`
	PublishTime   time.Time      `gorm:"column:publish_time" json:"publish_time"`
}

// TableName sets the insert table name for this struct type
func (model *VideoModel) TableName() string {
	return "video"
}

func AddVideoModel(m *VideoModel) error {
	return db.Save(m).Error
}

func DeleteVideoModelByID(id int) (bool, error) {
	if err := db.Delete(&VideoModel{}, id).Error; err != nil {
		return false, err
	}
	return db.RowsAffected > 0, nil
}

func DeleteVideoModel(condition string, args ...interface{}) (int64, error) {
	if err := db.Where(condition, args...).Delete(&VideoModel{}).Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func UpdateVideoModel(m *VideoModel) error {
	return db.Save(m).Error
}

func GetVideoModelByID(id int) (*VideoModel, error) {
	var m VideoModel
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetVideoModels(condition string, args ...interface{}) ([]*VideoModel, error) {
	res := make([]*VideoModel, 0)
	if err := db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
package model

import (
	"time"
)

type CommentModel struct {
	CommentID  int64     `gorm:"column:comment_id;primaryKey;unique;not null;autoIncrement" json:"comment_id"`
	VideoID    int64     `gorm:"column:video_id" json:"video_id"`
	UserID     int64     `gorm:"column:user_id" json:"user_id"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date"`
}

// TableName sets the insert table name for this struct type
func (model *CommentModel) TableName() string {
	return "comment"
}

// func AddCommentModel(m *CommentModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func DeleteCommentModelByID(id int) (bool, error) {
// 	if err := dao.MysqlDb.Delete(&CommentModel{}, id).Error; err != nil {
// 		return false, err
// 	}
// 	return dao.MysqlDb.RowsAffected > 0, nil
// }

// func DeleteCommentModel(condition string, args ...interface{}) (int64, error) {
// 	if err := dao.MysqlDb.Where(condition, args...).Delete(&CommentModel{}).Error; err != nil {
// 		return 0, err
// 	}
// 	return dao.MysqlDb.RowsAffected, nil
// }

// func UpdateCommentModel(m *CommentModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func GetCommentModelByID(id int) (*CommentModel, error) {
// 	var m CommentModel
// 	if err := dao.MysqlDb.First(&m, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &m, nil
// }

// func GetCommentModels(condition string, args ...interface{}) ([]*CommentModel, error) {
// 	res := make([]*CommentModel, 0)
// 	if err := dao.MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

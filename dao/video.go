package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"fmt"
)

type VideoDao interface {
	GetVideoModels() ([]*model.VideoModel, error)
	GetVideos() ([]model.VideoModel, error)
	AddVideoModel(m *model.VideoModel) error
	GetVideoModelByID(id int) (*model.VideoModel, error)
	GetVideoListByUserId(Userid int64) ([]model.VideoModel, error)
}

type videoDao struct{}

func NewVideoDao() VideoDao {
	return &videoDao{}
}

func (dao videoDao) GetVideos() ([]model.VideoModel, error) {
	// var idCardListEntity []IdCardEntity
	videos := make([]model.VideoModel, 0)
	if err := MysqlDb.Preload("Author").Find(&videos).Limit(30).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (dao videoDao) CheckFavorite(videoID int64, userID int64) (bool, error) {
	var res model.FavoriteModel
	if err := MysqlDb.Where("video_id = ? AND user_id = ?", videoID, userID).Find(&res).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (dao videoDao) CheckFollow(authorID int64, myID int64) (bool, error) {
	var res model.FollowModel
	if err := MysqlDb.Where("user_id = ? AND follwer_id = ?", authorID, myID).Find(&res).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (dao *videoDao) GetVideoModels() ([]*model.VideoModel, error) {
	res := make([]*model.VideoModel, 0)
	fmt.Println(MysqlDb)
	if err := MysqlDb.Table("video").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	fmt.Println(" GetVideoModels success")
	return res, nil
}

func (dao videoDao) AddVideoModel(m *model.VideoModel) error {
	return MysqlDb.Save(m).Error
}

func DeleteVideoModelByID(id int) (bool, error) {
	if err := MysqlDb.Delete(&model.VideoModel{}, id).Error; err != nil {
		return false, err
	}
	return MysqlDb.RowsAffected > 0, nil
}

func DeleteVideoModel(condition string, args ...interface{}) (int64, error) {
	if err := MysqlDb.Where(condition, args...).Delete(&model.VideoModel{}).Error; err != nil {
		return 0, err
	}
	return MysqlDb.RowsAffected, nil
}

func UpdateVideoModel(m *model.VideoModel) error {
	return MysqlDb.Save(m).Error
}

func (dao videoDao) GetVideoModelByID(id int) (*model.VideoModel, error) {
	var m model.VideoModel
	if err := MysqlDb.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao videoDao) GetVideoListByUserId(Userid int64) ([]model.VideoModel, error) {
	var videos []model.VideoModel
	if err := MysqlDb.Where("user_id = ?", Userid).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

package model

import (
	"database/sql"
)

type FavoriteModel struct {
	VideoID int           `gorm:"column:video_id;primaryKey;unique;not null" json:"video_id"`
	UserID  int           `gorm:"column:user_id;primaryKey;unique;not null" json:"user_id"`
	Status  sql.NullInt64 `gorm:"column:status" json:"status"`
}

// // TableName sets the insert table name for this struct type
func (FavoriteModel) TableName() string {
	return "favorite"
}


// func AddFavoriteModel(m *FavoriteModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func DeleteFavoriteModelByID(id int) (bool, error) {
// 	if err := dao.MysqlDb.Delete(&FavoriteModel{}, id).Error; err != nil {
// 		return false, err
// 	}
// 	return dao.MysqlDb.RowsAffected > 0, nil
// }

// func DeleteFavoriteModel(condition string, args ...interface{}) (int64, error) {
// 	if err := dao.MysqlDb.Where(condition, args...).Delete(&FavoriteModel{}).Error; err != nil {
// 		return 0, err
// 	}
// 	return dao.MysqlDb.RowsAffected, nil
// }

// func UpdateFavoriteModel(m *FavoriteModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func GetFavoriteModelByID(id int) (*FavoriteModel, error) {
// 	var m FavoriteModel
// 	if err := dao.MysqlDb.First(&m, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &m, nil
// }

// func GetFavoriteModels(condition string, args ...interface{}) ([]*FavoriteModel, error) {
// 	res := make([]*FavoriteModel, 0)
// 	if err := dao.MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
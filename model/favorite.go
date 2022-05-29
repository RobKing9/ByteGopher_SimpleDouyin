package model

import "database/sql"

type FavoriteModel struct {
	VideoID int           `gorm:"column:video_id;primaryKey;unique;not null" json:"video_id"`
	UserID  int           `gorm:"column:user_id;primaryKey;unique;not null" json:"user_id"`
	Status  sql.NullInt64 `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (model *FavoriteModel) TableName() string {
	return "favorite"
}

func AddFavoriteModel(m *FavoriteModel) error {
	return db.Save(m).Error
}

func DeleteFavoriteModelByID(id int) (bool, error) {
	if err := db.Delete(&FavoriteModel{}, id).Error; err != nil {
		return false, err
	}
	return db.RowsAffected > 0, nil
}

func DeleteFavoriteModel(condition string, args ...interface{}) (int64, error) {
	if err := db.Where(condition, args...).Delete(&FavoriteModel{}).Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func UpdateFavoriteModel(m *FavoriteModel) error {
	return db.Save(m).Error
}

func GetFavoriteModelByID(id int) (*FavoriteModel, error) {
	var m FavoriteModel
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetFavoriteModels(condition string, args ...interface{}) ([]*FavoriteModel, error) {
	res := make([]*FavoriteModel, 0)
	if err := db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
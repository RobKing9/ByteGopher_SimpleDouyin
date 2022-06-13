package model

type FollowModel struct {
	UserID    int `gorm:"column:user_id;primaryKey;unique;not null" json:"user_id"`
	FollwerID int `gorm:"column:follwer_id;primaryKey;unique;not null" json:"follwer_id"`
}

// // TableName sets the insert table name for this struct type
func (FollowModel) TableName() string {
	return "follow"
}

// func AddFollowModel(m *FollowModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func DeleteFollowModelByID(id int) (bool, error) {
// 	if err := dao.MysqlDb.Delete(&FollowModel{}, id).Error; err != nil {
// 		return false, err
// 	}
// 	return dao.MysqlDb.RowsAffected > 0, nil
// }

// func DeleteFollowModel(condition string, args ...interface{}) (int64, error) {
// 	if err := dao.MysqlDb.Where(condition, args...).Delete(&FollowModel{}).Error; err != nil {
// 		return 0, err
// 	}
// 	return dao.MysqlDb.RowsAffected, nil
// }

// func UpdateFollowModel(m *FollowModel) error {
// 	return dao.MysqlDb.Save(m).Error
// }

// func GetFollowModelByID(id int) (*FollowModel, error) {
// 	var m FollowModel
// 	if err := dao.MysqlDb.First(&m, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &m, nil
// }

// func GetFollowModels(condition string, args ...interface{}) ([]*FollowModel, error) {
// 	res := make([]*FollowModel, 0)
// 	if err := dao.MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
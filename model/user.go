package model

import "database/sql"

type UserModel struct {
	UserID        int            `gorm:"column:user_id;primaryKey;unique;not null;autoIncrement" json:"user_id"`
	UserName      sql.NullString `gorm:"column:user_name" json:"user_name"`
	Password      sql.NullString `gorm:"column:password" json:"password"`
	FollowCount   sql.NullInt64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount sql.NullInt64  `gorm:"column:follower_count" json:"follower_count"`
}

// TableName sets the insert table name for this struct type
func (model *UserModel) TableName() string {
	return "user"
}

func AddUserModel(m *UserModel) error {
	return db.Save(m).Error
}

func DeleteUserModelByID(id int) (bool, error) {
	if err := db.Delete(&UserModel{}, id).Error; err != nil {
		return false, err
	}
	return db.RowsAffected > 0, nil
}

func DeleteUserModel(condition string, args ...interface{}) (int64, error) {
	if err := db.Where(condition, args...).Delete(&UserModel{}).Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func UpdateUserModel(m *UserModel) error {
	return db.Save(m).Error
}

func GetUserModelByID(id int) (*UserModel, error) {
	var m UserModel
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserModels(condition string, args ...interface{}) ([]*UserModel, error) {
	res := make([]*UserModel, 0)
	if err := db.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
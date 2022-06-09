package dao

import (
	"ByteGopher_SimpleDouyin/model"
)

type UserDao interface {
	GetUserModelByID(id int) (*model.UserModel, error)
	GetUserByName(username string) (*model.UserModel, error)
	AddUserModel(m *model.UserModel) error
	GetCommonUserByID(id int) (*model.User, error)
}

type userDao struct{}

// TableName sets the insert table name for this struct type
// func (model *model.UserModel) TableName() string {
// 	return "user"
// }

func NewUserDao() UserDao {
	return &userDao{}
}

func (dao *userDao) GetUserModelByID(id int) (*model.UserModel, error) {
	var m model.UserModel
	if err := MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (dao *userDao) GetCommonUserByID(id int) (*model.User, error) {
	var m model.User
	if err := MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).Error; err != nil {
		return nil, err
	}
	if MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).RowsAffected == 0 {
		return nil, nil
	}
	return &m, nil
}

func (dao *userDao) GetUserByName(username string) (*model.UserModel, error) {
	var m model.UserModel

	if err := MysqlDb.Table("user").Where("user_name = ?", username).First(&m).Error; err != nil {
		print("哈哈哈1")
		return nil, err
	}
	return &m, nil
}

func (dao *userDao) AddUserModel(m *model.UserModel) error {
	return MysqlDb.Save(m).Error
}

/*
func AddUserModel(m *model.UserModel) error {
	return MysqlDb.Save(m).Error
}

func DeleteUserModelByID(id int) (bool, error) {
	if err := MysqlDb.Delete(&model.UserModel{}, id).Error; err != nil {
		return false, err
	}
	return MysqlDb.RowsAffected > 0, nil
}

func DeleteUserModel(condition string, args ...interface{}) (int64, error) {
	if err := MysqlDb.Where(condition, args...).Delete(&model.UserModel{}).Error; err != nil {
		return 0, err
	}
	return MysqlDb.RowsAffected, nil
}

func UpdateUserModel(m *model.UserModel) error {
	return MysqlDb.Save(m).Error
}


func GetUserModels(condition string, args ...interface{}) ([]*model.UserModel, error) {
	res := make([]*model.UserModel, 0)
	if err := MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

*/

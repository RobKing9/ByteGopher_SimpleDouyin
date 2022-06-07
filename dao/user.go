package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserDao interface {
	GetUserModelByID(id int) (*model.UserModel, error) 
}

type userDao struct{}

// TableName sets the insert table name for this struct type
// func (model *model.UserModel) TableName() string {
// 	return "user"
// }

func NewUserDao() UserDao {
	return &userDao{}
}

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

func (dao *userDao)GetUserModelByID(id int) (*model.UserModel, error) {
	var m model.UserModel
	if err := MysqlDb.Table("user").Where("user_id = ?", id).Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserModels(condition string, args ...interface{}) ([]*model.UserModel, error) {
	res := make([]*model.UserModel, 0)
	if err := MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetAllUserModels(db *gorm.DB) ([]*model.UserModel, error) {
	fmt.Println("before make")
	res := make([]*model.UserModel, 0)
	fmt.Println("after make")
	if err := db.Find(&res).Error; err != nil {
		fmt.Println("Find Error!")
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

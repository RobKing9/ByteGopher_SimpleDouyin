package model

import (
	"ByteGopher_SimpleDouyin/dao"
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
)

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
	return dao.MysqlDb.Save(m).Error
}

func DeleteUserModelByID(id int) (bool, error) {
	if err := dao.MysqlDb.Delete(&UserModel{}, id).Error; err != nil {
		return false, err
	}
	return dao.MysqlDb.RowsAffected > 0, nil
}

func DeleteUserModel(condition string, args ...interface{}) (int64, error) {
	if err := dao.MysqlDb.Where(condition, args...).Delete(&UserModel{}).Error; err != nil {
		return 0, err
	}
	return dao.MysqlDb.RowsAffected, nil
}

func UpdateUserModel(m *UserModel) error {
	return dao.MysqlDb.Save(m).Error
}

func GetUserModelByID(db *gorm.DB, id int) (*UserModel, error) {
	var m UserModel
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func GetUserModels(condition string, args ...interface{}) ([]*UserModel, error) {
	res := make([]*UserModel, 0)
	if err := dao.MysqlDb.Where(condition, args...).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetAllUserModels(db *gorm.DB) ([]*UserModel, error) {
	fmt.Println("before make")
	res := make([]*UserModel, 0)
	fmt.Println("after make")
	if err := db.Find(&res).Error; err != nil {
		fmt.Println("Find Error!")
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
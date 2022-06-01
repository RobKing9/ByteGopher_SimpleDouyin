package dao

import "ByteGopher_SimpleDouyin/model"

type UserModel struct {
	model.User	`gorm:"embedded"`
}

func (u UserModel) TableName() string {
	return "user"
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (u *UserModel) Save() error {
	return nil
}

func (u *UserModel) SearchUser(user, pwd string) (model.User,error) {
	usr := model.User{}

	var _ = pwd

	tx := GetDB().Where("user_name = ?",user).Table("user").Find(&usr)
	return usr,tx.Error
}

func (u *UserModel) SearchUserById(userId int64) (model.User,error) {
	usr := model.User{}



	tx := GetDB().Table("user").Where("user_id = ?",userId).Find(&usr)
	//log.Println(userId)
	//log.Printf("%#v",usr)
	return usr,tx.Error
}

func (u *UserModel) Search() {

}


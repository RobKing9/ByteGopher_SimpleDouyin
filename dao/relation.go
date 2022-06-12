package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"encoding/json"
	"fmt"
	"log"
)

// 关系数据库表设计
// 关注列表 follow: id | user_id | follow (存储 to_suer_id)
// 粉丝列表 fans:   id | user_id | fans	(存储 to_user_id)


// Fans 结构模型
type Fans struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	int64	`json:"to_user_id" gorm:"column:fans"`
}

// Follow 结构模型
type Follow struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	int64	`json:"to_user_id" gorm:"column:follow"`
}

const (
	FansListType	= "fans"
	FollowListType	= "follow"
)


//// GetFollowByUserId 通过 userId 获取关注列表
//func GetFollowByUserId(userid int64) []model.UserModel {
//	return GetRelationByUserId(userid,FollowListType)
//}
//
//// GetFansByUserId 通过 userId 获取粉丝列表
//func GetFansByUserId(userid int64) []model.UserModel {
//	return GetRelationByUserId(userid,FansListType)
//}

// GetRelationByUserId 得到特定的用户列表切片
func GetRelationByUserId(userid int64,listType string) []model.UserModel {
	switch listType {
	case FollowListType:
		return getUserList(userid,FollowListType)
	case FansListType:
		return getUserList(userid,FansListType)
	}
	return []model.UserModel{}
}


// 得到用户结构体切片
func getUserList(userid int64,listType string) []model.UserModel {
	users := []model.UserModel{}

	queryStr := fmt.Sprintf("%v.%v",listType,listType)
	subQuery := GetDB().Select(queryStr).Where("user_id = ?",userid).Table(listType)
	if subQuery.Error != nil {
		// TODO: 写入日志？
		log.Println(subQuery.Error)
	}

	tx := GetDB().Table("user").Select("*").Where("user_id in (?)",subQuery).Find(&users)
	if tx.Error != nil {
		// TODO: 写入日志？
		log.Println(tx.Error)
	}

	return users
}



// GetUserJson 返回 某个 userJSON
func GetUserJson(userid int64, isfollow bool) string {
	info, err := getUserInfo(userid)
	if err != nil {
		return ""
	}

	user := model.User{
		Id: info.UserID,
		Name: info.UserName,
		FollowCount: info.FollowCount,
		FollowerCount: info.FollowerCount,
		IsFollow: isfollow,
	}

	infoJson, err := json.Marshal(user)
	if err != nil {
		return ""
	}

	return string(infoJson)
}



func getUserInfo(userid int64) (model.UserModel,error) {
	var m model.UserModel
	tx := MysqlDb.Table("user").Where("user_id = ?", userid).Find(&m)

	return m, tx.Error
}



// GetFans 查询粉丝表
func GetFans(userid int64) ([]Fans,error) {
	fans := []Fans{}
	tx := GetDB().Table("fans").Where("user_id = ?",userid).Find(&fans)

	return fans,tx.Error
}



// GetFollow 查询关注表
func GetFollow(userid int64) ([]Follow,error) {
	fans := []Follow{}
	tx := GetDB().Table("follow").Where("user_id = ?",userid).Find(&fans)

	return fans,tx.Error
}



// SaveFollowCount 更新关注数
func SaveFollowCount(uid, followCount int64) error {
	user := model.UserModel{
		UserID: uid,
	}

	tx := GetDB().Model(&user).Table("user").Select("follow_count").Update("follow_count",followCount)
	return tx.Error
}



// SaveFansCount 更新粉丝数
func SaveFansCount(uid, fansCount int64) error {
	user := model.UserModel{
		UserID: uid,
	}

	tx := GetDB().Model(&user).Table("user").Select("follow_count").Update("follower_count",fansCount)
	return tx.Error
}
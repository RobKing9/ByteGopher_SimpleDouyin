package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"encoding/json"
)

// 关系数据库表设计
// 关注列表 follow: id | user_id | followinfo (存储 to_suer_id JSON)
// 粉丝列表 fans:   id | user_id | fansinfo	(存储 to_user_id JSON) | is_follow


// FansModel 结构模型
type FansModel struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	string	`json:"fans" gorm:"column:fansinfo"`
	IsFollow	bool	`json:"is_follow" gorm:"column:is_follow"`
}

// FollowModel 结构模型
type FollowModel struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	string	`json:"follow" gorm:"column:followinfo"`
}



// SaveFollowInToTable 把自己加入对方粉丝表，和自己的关注表
func SaveFollowInToTable(myId, toUserid int64) error {
	user, err := getUser(myId)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fan := FansModel{
		UserId: toUserid,
		ToUserId: string(marshal),
	}

	tx := GetDB().Table("fanstable").Save(&fan)
	if tx.Error != nil {
		return tx.Error
	}

	return addMyFollow(myId,toUserid)
}



// DeleteFansInToTable 把自己从对方粉丝表中删除，和自己的关注表
func DeleteFansInToTable(myId, toUserid int64) error {
	user, err := getUser(myId)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fan := FansModel{
		UserId: toUserid,
		ToUserId: string(marshal),
	}

	tx := GetDB().Table("fanstable").Delete(&fan)
	if tx.Error != nil {
		return tx.Error
	}

	return deleteMyFollow(myId,toUserid)
}




// getUser 从 user 表得到用户信息
func getUser(uid int64) (model.User,error) {
	uModel := model.UserModel{}

	tx := GetDB().Table("user").Where("user_id = ?",uid).Find(&uModel)

	u := model.User{
		Id: uModel.UserID,
		Name: uModel.UserName,
		FollowCount: uModel.FollowCount,
		FollowerCount: uModel.FollowerCount,
	}

	return u,tx.Error
}


// 把对方添加进自己的关注表
func addMyFollow(myId,toUserId int64) error {
	toUser, err := getUser(toUserId)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(toUser)
	if err != nil {
		return err
	}

	follow := FollowModel{
		UserId: myId,
		ToUserId: string(marshal),
	}

	tx := GetDB().Table("followtable").Save(&follow)

	return tx.Error
}


// 自己的关注表删除对方
func deleteMyFollow(myId,toUserId int64) error {
	toUser, err := getUser(toUserId)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(toUser)
	if err != nil {
		return err
	}

	follow := FollowModel{
		UserId: myId,
		ToUserId: string(marshal),
	}

	tx := GetDB().Table("followtable").Delete(&follow)

	return tx.Error
}


// GetFollowList 从 followtable 表中获取关注列表
func GetFollowList(uid int64) ([]model.User,error) {
	f := []FollowModel{}
	u := []model.User{}

	tx := GetDB().Table("followtable").Where("user_id = ?",uid).Find(&f)

	for _,v := range f {
		user := model.User{}

		json.Unmarshal([]byte(v.ToUserId), &user)

		user.IsFollow = true
		u = append(u,user)
	}

	err := countUsers(uid)
	if err != nil {
		return nil, err
	}

	return u,tx.Error
}



// GetFansList 从 fanstable 表中获取粉丝列表
func GetFansList(uid int64) ([]model.User,error) {
	f := []FansModel{}
	u := []model.User{}
	fMap := make(map[int64]bool)

	// 获取关注表
	flist,err := GetFollowList(uid)
	if err != nil {
		return u,err
	}
	for _,v := range flist {
		if !fMap[v.Id] {
			fMap[v.Id] = true
		}
	}

	tx := GetDB().Table("fanstable").Where("user_id = ?",uid).Find(&f)

	for _,v := range f {
		user := model.User{}

		json.Unmarshal([]byte(v.ToUserId), &user)

		if fMap[v.UserId] {
			v.IsFollow = true
		}

		u = append(u,user)
	}

	return u,tx.Error
}


// CountUsers 统计关注数和粉丝数
// 并更改
func countUsers(uid int64) error {
	var followCount int64
	tx := GetDB().Table("followtable").Where("user_id = ?",uid).Count(&followCount)
	if tx.Error != nil {
		return tx.Error
	}

	var fansCount int64
	tx = GetDB().Table("fanstable").Where("user_id = ?",uid).Count(&fansCount)
	if tx.Error != nil {
		return tx.Error
	}

	u := model.UserModel{
		UserID: uid,
	}

	tx = GetDB().Model(&u).Table("user").Updates(model.UserModel{FollowerCount: fansCount,FollowCount: followCount})
	return tx.Error
}

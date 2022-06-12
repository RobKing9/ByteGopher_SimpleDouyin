package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"encoding/json"
	"log"
)

// 关系数据库表设计
// 关注列表 follow: id | user_id | to_user_id | data (存储 to_suer JSON)
// 粉丝列表 fans:   id | user_id | to_user_id | data (存储 to_user JSON)


// FansModel 结构模型
type FansModel struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	int64	`json:"to_user_id" gorm:"column:to_user_id"`
	Data		string	`json:"data" gorm:"column:data"`
	IsFollow	bool	`json:"is_follow" gorm:"-"`
}

func (f FansModel) TableName() string {
	return "fanstable"
}

// FollowModel 结构模型
type FollowModel struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	int64	`json:"to_user_id" gorm:"column:to_user_id"`
	Data		string	`json:"data" gorm:"column:data"`
}

func (f Follow) TableName() string {
	return "followtable"
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
		ToUserId: myId,
		Data: string(marshal),
	}

	tx := GetDB().Table("fanstable").Save(&fan)
	if tx.Error != nil {
		return tx.Error
	}

	return addMyFollow(myId,toUserid)
}



// DeleteFansInToTable 把自己从对方粉丝表中删除，和自己的关注表
func DeleteFansInToTable(myId, toUserid int64) error {
	//user, err := getUser(myId)
	//if err != nil {
	//	return err
	//}
	//
	//marshal, err := json.Marshal(user)
	//if err != nil {
	//	return err
	//}

	fan := FansModel{
		UserId: toUserid,
		ToUserId: myId,
	}

	log.Printf("%#v",fan)
	tx := GetDB().Exec("delete from fanstable where user_id = ? and to_user_id = ?",toUserid,myId)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	err := deleteMyFollow(myId,toUserid)
	if err != nil {
		return err
	}

	return nil
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
		ToUserId: toUserId,
		Data: string(marshal),
	}

	tx := GetDB().Table("followtable").Save(&follow)

	return tx.Error
}


// 自己的关注表删除对方
func deleteMyFollow(myId,toUserId int64) error {
	//toUser, err := getUser(toUserId)
	//if err != nil {
	//	return err
	//}
	//
	//marshal, err := json.Marshal(toUser)
	//if err != nil {
	//	return err
	//}

	//follow := FollowModel{
	//	UserId: myId,
	//	ToUserId: toUserId,
	//}


	tx := GetDB().Exec("delete from followtable where user_id = ? and to_user_id = ?",myId,toUserId)


	return tx.Error
}


// GetFollowList 从 followtable 表中获取关注列表
func GetFollowList(uid int64) ([]model.User,error) {
	f := []FollowModel{}
	u := []model.User{}

	tx := GetDB().Table("followtable").Where("user_id = ?",uid).Find(&f)

	for _,v := range f {
		user := model.User{}

		json.Unmarshal([]byte(v.Data), &user)

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

	f := []FansModel{}
	tx := GetDB().Table("fanstable").Where("user_id = ?",uid).Find(&f)

	for _,v := range f {
		user := model.User{}

		json.Unmarshal([]byte(v.Data), &user)

		if fMap[v.ToUserId] {
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

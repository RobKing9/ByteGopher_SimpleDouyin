package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

// 关系数据库表设计
// 关注列表 follow: id | user_id | to_user_id | data (存储 to_suer JSON) | is_follow
// 粉丝列表 fans:   id | user_id | to_user_id | data (存储 to_user JSON) | is_follow

//drop table if EXISTS followtable;
//
//create table followtable (
//id int PRIMARY key AUTO_INCREMENT,
//user_id BIGINT(20) not null,
//to_user_id BIGINT(20)  not null,
//data BLOB(255),
//is_follow tinyint(1) not null
//);

var Wg sync.WaitGroup

// FansModel 结构模型
type FansModel struct {
	Id 			int64	`json:"-" gorm:"column:id"`
	UserId 		int64	`json:"user_id" gorm:"column:user_id"`
	ToUserId 	int64	`json:"to_user_id" gorm:"column:to_user_id"`
	Data		string	`json:"data" gorm:"column:data"`
	IsFollow	bool	`json:"is_follow" gorm:"column:is_follow"`
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
	IsFollow	bool	`json:"is_follow" gorm:"column:is_follow"`
}

func (f FollowModel) TableName() string {
	return "followtable"
}


// SaveFollowInToTable 把自己加入对方粉丝表，和自己的关注表
func SaveFollowInToTable(myId, toUserid int64) error {
	if myId == toUserid {
		return errors.New("can't follow yourself")
	}

	// 判断即将插入的数据是否存在
	b, err := judgeExist(toUserid, myId, fans)
	if err != nil || b == true {
		return err
	}

	user, err := getUser(myId)
	if err != nil {
		return err
	}


	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}

	fanChan := make(chan FansModel,1)


	go func(myId, toUserid int64, str string ,fanChan chan FansModel) {
		var exist bool
		// 判断自己是否在对方关注表中
		flist,err := GetFollowList(toUserid)
		if err != nil {
			log.Println(err)
		}
		for _,v := range flist {
			if v.Id == myId {
				exist = true
			}
		}

		fan := FansModel{
			UserId: toUserid,
			ToUserId: myId,
			Data: str,
			IsFollow: exist,
		}

		fanChan <- fan
		close(fanChan)

	}(myId, toUserid,string(marshal),fanChan)


	Wg.Add(2)
	defer Wg.Wait()
	go func(myId, toUserid int64, fanChan chan FansModel) {
		fan := <-fanChan

		tx := GetDB().Table("fanstable").Save(&fan)
		if tx.Error != nil {
			log.Println(tx.Error)
		}

		// 更改自己粉丝列表状态
		tx = GetDB().Exec("update fanstable set is_follow = ? where user_id = ? and to_user_id = ?",true,myId,toUserid)
		if tx.Error != nil {
			log.Println(tx.Error)
		}

		Wg.Done()
	}(myId, toUserid,fanChan)

	go addMyFollow(myId,toUserid)

	return nil
}



// DeleteFansInToTable 把自己从对方粉丝表中删除，和自己的关注表
func DeleteFansInToTable(myId, toUserid int64) error {

	Wg.Add(1)
	defer Wg.Wait()

	go func(myId, toUserid int64) {
		tx := GetDB().Exec("update fanstable set is_follow = ? where user_id = ? and to_user_id = ?",false,myId,toUserid)
		if tx.Error != nil {
			log.Println(tx.Error)
		}


		tx = GetDB().Exec("delete from fanstable where user_id = ? and to_user_id = ?",toUserid,myId)
		if tx.Error != nil {
			log.Println(tx.Error)
		}

		Wg.Done()
	}(myId, toUserid)


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
func addMyFollow(myId,toUserId int64)  {

	// 判断即将插入的数据是否存在
	b, err := judgeExist(myId, toUserId, follow)
	if err != nil || b == true {
		log.Println(err)
	}

	toUser, err := getUser(toUserId)
	if err != nil {
		log.Println(err)
	}

	marshal, err := json.Marshal(toUser)
	if err != nil {
		log.Println(err)
	}

	follow := FollowModel{
		UserId: myId,
		ToUserId: toUserId,
		Data: string(marshal),
		IsFollow: true,
	}

	tx := GetDB().Table("followtable").Save(&follow)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	Wg.Done()
}


// 自己的关注表删除对方
func deleteMyFollow(myId,toUserId int64) error {

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

		user.IsFollow = v.IsFollow
		u = append(u,user)
	}

	return u,tx.Error
}



// GetFansList 从 fanstable 表中获取粉丝列表
func GetFansList(uid int64) ([]model.User,error) {
	u := []model.User{}

	f := []FansModel{}
	tx := GetDB().Table("fanstable").Where("user_id = ?",uid).Find(&f)

	for _,v := range f {
		user := model.User{}

		json.Unmarshal([]byte(v.Data), &user)

		user.IsFollow = v.IsFollow

		u = append(u,user)
	}

	return u,tx.Error
}


// CountUsers 统计关注数和粉丝数
// 并更改
func CountUsers(uid int64) {
	var followCount int64
	tx := GetDB().Table("followtable").Where("user_id = ?",uid).Count(&followCount)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	var fansCount int64
	tx = GetDB().Table("fanstable").Where("user_id = ?",uid).Count(&fansCount)
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	u := model.UserModel{
		UserID: uid,
	}

	tx = GetDB().Model(&u).Table("user").Updates(model.UserModel{FollowerCount: fansCount,FollowCount: followCount})
	if tx.Error != nil {
		log.Println(tx.Error)
	}

	Wg.Done()
}


const (
	fans	= "fanstable"
	follow	= "followtable"
)



// 判断插入的数据是否存在
// tableType 有 fanstable / followtable 两种
func judgeExist(uid, toUserId int64,tableType string) (bool,error) {
	switch tableType {
	case fans:
		fan := FansModel{}
		tx := GetDB().Table(tableType).Where("user_id = ? and to_user_id = ?",uid,toUserId).Find(&fan)
		return fan.Id > 0,tx.Error
	case follow:
		follow := FollowModel{}
		tx := GetDB().Table(tableType).Where("user_id = ? and to_user_id = ?",uid,toUserId).Find(&follow)
		return follow.Id > 0,tx.Error
	}
	return true,nil
}
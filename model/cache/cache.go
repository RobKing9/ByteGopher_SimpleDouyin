package cache

/*若无 redis 临时测试用 */

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"sync"
)

type UserIdType int64

// 充当缓存( 每个用户都需要维护自己的列表缓存 )
var mu sync.RWMutex
var followList = make(map[UserIdType][]model.User)
var fansList = make(map[UserIdType][]model.User)


func GetTestData(userid int64) {
	followList[UserIdType(userid)] = []model.User{
		{3, "aya3", "", 0, 0, false},
	}
	fansList[UserIdType(userid)] = []model.User{
		{2, "aya2", "", 0, 0, false},
	}
}

// CreateCacheList 当用户注册成功后就维护该列表缓存
// 其中数据需要持久化
func CreateCacheList(userid UserIdType) {
	// 判断用户的列表是否已经存在
	followList[userid] = nil
	fansList[userid] = nil
}

// GetFollowList 获取关注列表
func GetFollowList(userid UserIdType)  []model.User {
	mu.Lock()
	defer mu.Unlock()
	follows := followList[userid]

	return follows
}

// GetFansList 获取粉丝列表
func GetFansList(userid UserIdType) []model.User {
	mu.Lock()
	defer mu.Unlock()

	fans := fansList[userid]
	return fans
}

// Follow 关注操作
func Follow(toUserId,userId int64,actionType int32) error {
	mu.Lock()
	defer mu.Unlock()

	// 获取自身信息
	user,err := dao.NewUserModel().SearchUserById(userId)
	if err != nil {
		return err
	}

	switch actionType {
	case 1:
		toFans := fansList[UserIdType(toUserId)]
		toFans = append(toFans,user)
	case 2:
		toFans := fansList[UserIdType(toUserId)]

		// 查找自身 Id ，从粉丝列表中删除对应 结构体
		for i,v := range toFans{
			if v.UserId == user.UserId {
				toFans = append(toFans[:i],toFans[i+1:]...)
				break
			}
		}
	}

	return nil
}

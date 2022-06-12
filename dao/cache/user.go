package cache

import (
	"ByteGopher_SimpleDouyin/dao"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// GetUserInfoJson 直接从 redis 缓存中获取用户信息
func GetUserInfoJson(userid int64) string {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()


	key := fmt.Sprintf("user:%v",userid)
	result, err := GetRdb().Get(ctx,key).Result()
	if err != nil {
		if err == redis.Nil {
			// 若 redis 返回为空就从数据库中获取
			// 但得到的用户数据必须为 userInfo 的 JSON 形式
			info := dao.GetUserJson(userid, false)

			// TODO: 重新写入 redis
			return info
		}

		// 是否需要写入日志文件
		log.Println("GetUserInfoJson: ",err)
		return ""
	}

	return result
}


// SetUserInfoJson 设置 redis 缓存用户信息
func SetUserInfoJson(userid int64) {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()


	// 查看缓存中是否存在
	key := fmt.Sprintf("user:%v",userid)

	result, err := GetRdb().Get(ctx,key).Result()
	if result != "0" || err != nil {
		return
	}

	// 若不存在，设置缓存
	info := dao.GetUserJson(userid, true)
	result, err = GetRdb().Set(ctx,key,info,time.Hour).Result()
	if result != "0" || err != nil {
		return
	}
}


// GetUserStruct  解析用户结构体
// 直接从 redis 缓存中获取用户信息，并转为 model.User 结构体
//func GetUserStruct(userid int64) model.User {
//
//	str := GetUserInfoJson(userid)
//	if str == "" {
//		return model.User{UserId: -1}
//	}
//
//	user := model.User{}
//
//	// 反序列化，得到user结构体
//	err := json.Unmarshal([]byte(str),&user)
//	if err != nil {
//		// 是否需要写入日志文件
//		log.Println(err)
//	}
//
//	return user
//}
package cache

import (
	"ByteGopher_SimpleDouyin/dao"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"

	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
)

// 此处函数直接操作 redis

const (
	fansList 	= "fans"
	followList	= "follow"
)

// GetFansListByUserId 通过 userId 查询 粉丝列表
func GetFansListByUserId(userid int64) []model.User{
	return getUserListByUserId(userid,fansList)
}

// GetFollowListByUserId 通过 userId 查询 关注列表
func GetFollowListByUserId(userid int64) []model.User{
	return getUserListByUserId(userid,followList)
}


// 通过 userId 获得用户表
func getUserListByUserId(userid int64, listType string) []model.User {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	// inChan 在协程传递 userInfo 的JSON 值
	// userInfo 的 JSON 值，如：
	// {"id":1,"name":"aya","follow_count":0,"follower_count":0,"is_follow":false}
	// 此JSON值会作为 redis 缓存中 user:id（key）的 value 值
	inChan := make(chan string,5)

	// 传递 粉丝列表
	outChan := make(chan []model.User,1)


	key := fmt.Sprintf("%v:%v",listType,userid)
	// 获取粉丝id列表
	usersId,err := GetRdb().ZRange(ctx,key,0,-1).Result()
	if err != nil {
		if err == redis.Nil {
			// TODO: 配置策略更新缓存
			// 若 redis 返回为空, 需要先重新设置缓存
			// 再重新执行本次操作
			users := dao.GetRelationByUserId(userid,listType)

			RdbWg.Add(1)
			go changeStruct(userid,users,outChan,listType)
			RdbWg.Wait()

			return  <- outChan
		}
		// 错误是否需要写入日志文件?
		log.Println("GetFansByUserId: ",err)
	}


	//log.Println("getUserListByUserId",usersId)

	RdbWg.Add(1)
	go getUsers(usersId,inChan)     // 通过 userid 获取userInfo
	go jsonToStruct(inChan,outChan) // 得到 粉丝列表
	RdbWg.Wait()


	return <-outChan
}


// 结构体转化
// TODO: 转换结构体 model.UserModle -> model.User
func changeStruct(uid int64,users []model.UserModel,outChan chan []model.User,listType string) {
	u := []model.User{}


	// TODO: 更换逻辑
	if listType == followList {
		for _, v := range users {
			user := model.User{
				Id: v.UserID,
				Name: v.UserName,
				FollowCount: v.FollowCount,
				FollowerCount: v.FollowerCount,
				IsFollow: true,
			}

			u = append(u,user)
		}
	} else {
		userMap := make(map[int64]bool)

		follow ,_ := dao.GetFollow(uid)
		for _,v := range follow {
			if !userMap[v.UserId] {
				userMap[v.UserId] = true
			}
		}

		for _, v := range users {
			user := model.User{
				Id: v.UserID,
				Name: v.UserName,
				FollowCount: v.FollowCount,
				FollowerCount: v.FollowerCount,
				IsFollow: userMap[v.UserID],
			}

			u = append(u,user)
		}
	}


	outChan <- u
	close(outChan)
	RdbWg.Done()
}


// getUsers 通过 userid 获取userInfo
func getUsers(userIds []string, inChan chan string) {
	for _,v := range userIds {
		id, _ := strconv.ParseInt(v,0,64)
		infoJson := GetUserInfoJson(id)		// 从用户信息表中获取userInfo
		inChan <- infoJson
	}

	close(inChan)
}

// jsonToStruct 得到 用户列表切片
// 将 json 形式的 user 数据转化为 model.User 结构体
func jsonToStruct(inChan chan string, outChan chan []model.User) {
	fanslist := []model.User{}

	for  {
		user := model.User{}

		str,ok := <-inChan
		if !ok {	// ok 为 false 时 chan 关闭
			break
		}

		// 反序列化，得到user结构体
		err := json.Unmarshal([]byte(str),&user)
		if err != nil {
			log.Println(err)
		}

		fanslist = append(fanslist,user)
	}


	outChan <- fanslist

	close(outChan)
	RdbWg.Done()
}



// FollowAction 关注
// 在对方粉丝列表加入自己的 userid
// 将当前时间戳作为 score
func FollowAction(toUserId,myId int64) error {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	key1 := fmt.Sprintf("fans:%v",toUserId)
	key2 := fmt.Sprintf("follow:%v",myId)

	score := utils.NowTimeToFloat64()
	if score == -1 {
		return errors.New("illegal operation")
	}

	// 在自己 follow 中加入对方 id
	_, err := GetRdb().ZAdd(ctx,key2, &redis.Z{score, toUserId}).Result()
	if err != nil {
		if err == redis.Nil {
			// 若 redis 返回为空, 需要先重新设置缓存
			// 再重新执行本次操作
			err := zSetAddFan(myId,followList)
			if err != nil {
				log.Println(err)
				return err
			}
		}
		log.Println(err)
		return err
	}

	// 在对方 fans 中加入自己 id
	_, err = GetRdb().ZAdd(ctx,key1, &redis.Z{score, myId}).Result()
	if err != nil {
		if err == redis.Nil {
			// 若 redis 返回为空, 需要先重新设置缓存
			// 再重新执行本次操作
			err := zSetAddFan(toUserId,fansList)
			if err != nil {
				log.Println(err)
				return err
			}
		}
		log.Println(err)
		return err
	}

	SetUserInfoJson(toUserId)
	SetUserInfoJson(myId)

	// 统计关注数
	countFollow(myId)

	return nil
}


// countFollow 统计关注数
func countFollow(uid int64) {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	key := fmt.Sprintf("follow:%v",uid)
	result ,err := GetRdb().ZCard(ctx,key).Result()
	if result == 0 || err != nil {
		return
	}

	err = dao.SaveFollowCount(uid,result)
	if err != nil {
		log.Println("countFollow: ",err)
	}
}



// CancelFollowAction 取消关注
// 在对方粉丝列表删除自己的 userid
func CancelFollowAction(toUserId,myId int64) error {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	key1 := fmt.Sprintf("fans:%v",toUserId)
	key2 := fmt.Sprintf("follow:%v",toUserId)

	// 把自己从对方 fans 中删除
	_,err := GetRdb().ZRem(ctx,key1,myId).Result()
	if err != nil {
		if err == redis.Nil {
			// 若 redis 返回为空, 需要先重新设置缓存
			// 再重新执行本次操作
			err := zSetAddFan(toUserId,fansList)
			if err != nil {
				log.Println(err)
				return err
			}
		}
		log.Println(err)
		return err
	}

	// 把对方从自己 follow 中删除
	_,err = GetRdb().ZRem(ctx,key2,toUserId).Result()
	if err != nil {
		if err == redis.Nil {
			// 若 redis 返回为空, 需要先重新设置缓存
			// 再重新执行本次操作
			err := zSetAddFan(myId,followList)
			if err != nil {
				log.Println(err)
				return err
			}
		}
		log.Println(err)
		return err
	}

	// 统计粉丝数
	countFans(myId)

	return nil
}


// countFans 统计关注数
func countFans(uid int64) {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	key := fmt.Sprintf("fans:%v",uid)
	result ,err := GetRdb().ZCard(ctx,key).Result()
	if result == 0 || err != nil {
		return
	}

	err = dao.SaveFansCount(uid,result)
	if err != nil {
		log.Println("countFans: ",err)
	}
}



// zSetAddFan 在缓存 fans:userid 中添加成员
// TODO: 如何恢复缓存?
func zSetAddFan(userId int64,listType string) error {
	// 建立 context 上下文，用于 redis 的操作
	// ctx 在 3 秒后失效, 或者在函数返回后取消
	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
	defer cancel()

	key := fmt.Sprintf("%v:%v",listType,userId)

	fans, err := dao.GetFans(userId)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range fans {
		score := utils.NowTimeToFloat64()
		if score == -1 {
			return errors.New("illegal operation")
		}

		_, err = GetRdb().ZAdd(ctx,key, &redis.Z{score, v.ToUserId}).Result()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
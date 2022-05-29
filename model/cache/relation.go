package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis"

	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils/timeTool"
)

// GetFansListByUserId 通过 userId 查询 粉丝列表
func GetFansListByUserId(userid int64) []model.User{
	return getUserListByUserId(userid,model.FansList)
}

// GetFollowListByUserId 通过 userId 查询 关注列表
func GetFollowListByUserId(userid int64) []model.User{
	return getUserListByUserId(userid,model.FollowList)
}



func getUserListByUserId(userid int64, listType string) []model.User {

	// 在协程传递 userInfo 的JSON 值
	inChan := make(chan string,5)

	// 传递 粉丝列表
	outChan := make(chan []model.User,1)

	str := fmt.Sprintf("%v:%v",listType,userid)

	// 获取粉丝id列表
	usersId,err := GetRdb().ZRange(str,0,-1).Result()
	if err != nil {
		// 是否需要写入日志文件
		log.Println("GetFansByUserId: ",err)
	}


	RdbWg.Add(1)
	go getFans(usersId,inChan)		// 通过 userid 获取userInfo
	go jsonToStruct(inChan,outChan)	// 得到 粉丝列表
	RdbWg.Wait()


	return <-outChan
}



// getFans 通过 userid 获取userInfo
func getFans(userIds []string, inChan chan string) {
	for _,v := range userIds {
		id, _ := strconv.ParseInt(v,0,64)
		infoJson := GetUserInfoJson(id)		// 从用户信息表中获取userInfo
		inChan <- infoJson
	}

	close(inChan)

}

// jsonToStruct 得到 用户列表
func jsonToStruct(inChan chan string, outChan chan []model.User) {
	fanslist := []model.User{}

	for  {
		user := model.User{}

		str,ok := <-inChan
		if !ok {
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
	key := fmt.Sprintf("fans:%v",toUserId)

	score := timeTool.NowTimeToFloat64()
	if score == -1 {
		return errors.New("illegal operation")
	}

	_, err := GetRdb().ZAdd(key, redis.Z{score, myId}).Result()
	if err != nil {
		return err
	}
	return nil
}

// CancelFollowAction 取消关注
// 在对方粉丝列表删除自己的 userid
func CancelFollowAction(toUserId,myId int64) error {
	key := fmt.Sprintf("fans:%v",toUserId)

	_,err := GetRdb().ZRem(key,myId).Result()
	if err != nil {
		return err
	}
	return nil
}

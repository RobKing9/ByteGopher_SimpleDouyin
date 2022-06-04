package cache


import (
	"encoding/json"
	"fmt"
	"log"

	"ByteGopher_SimpleDouyin/model"
)

func GetUserInfoJson(userid int64) string {

	key := fmt.Sprintf("user:%v",userid)
	result, err := GetRdb().Get(key).Result()
	if err != nil {
		// 是否需要写入日志文件
		log.Println("GetUserInfoJson: ",err)
		return ""
	}

	return result
}

// JsonToStruct  解析用户结构体
func JsonToStruct(userid int64) model.User {

	str := GetUserInfoJson(userid)
	if str == "" {
		return model.User{UserId: -1}
	}

	user := model.User{}

	// 反序列化，得到user结构体
	err := json.Unmarshal([]byte(str),&user)
	if err != nil {
		// 是否需要写入日志文件
		log.Println(err)
	}

	return user
}
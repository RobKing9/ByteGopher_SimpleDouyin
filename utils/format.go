package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

func MapToJson(data interface{}) string {
    byteStr, _ := json.Marshal(data)
    return string(byteStr)
}


func UnixToTime(unix string) string {
	log.Println("未转化时间戳", unix)
	unixInt64, _ := strconv.ParseInt(unix, 10, 64)
	timeTemplate := "2006-01-02 15:04:05" //常规类型
	return time.Unix(unixInt64, 0).Format(timeTemplate)
}
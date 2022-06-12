package cache

import (
	"ByteGopher_SimpleDouyin/dao"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	//rdb *redis.Client

	// RdbWg 全局等待组
	RdbWg sync.WaitGroup
)

func GetRdb() *redis.Client {
	return dao.GetRdbClient()
}

//func InitRDB() error {
//	// 建立 context 上下文，用于 redis 的操作
//	// ctx 在 3 秒后失效, 或者在函数返回后取消
//	ctx, cancel := context.WithTimeout(context.Background(),time.Second*3)
//	defer cancel()
//
//	// 连接 redis 客户端
//	db := redis.NewClient(&redis.Options{
//		Addr: "121.40.120.222:46379"/*在配置文件中重写*/,
//		Password: "Byt3G0pheR51522zzwlwlbb"/*在配置文件中重写*/,
//		DB: 0,
//	})
//
//	// 测试连接
//	ping, err := db.Ping(ctx).Result()
//	if err != nil {
//		return err
//	}
//	if ping != "PONG" {
//		return errors.New("connect failed")
//	}
//
//	// TODO: 连接池
//
//	rdb = db
//	return nil
//}
//

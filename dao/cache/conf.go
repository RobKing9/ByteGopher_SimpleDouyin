package cache


import (
	"errors"
	"sync"

	"github.com/go-redis/redis"
)

var rdb *redis.Client
var RdbWg sync.WaitGroup

func GetRdb() *redis.Client {
	return rdb
}

func InitRDB() error {
	db := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "123456",
		DB: 0,
	})

	ping, err := db.Ping().Result()
	if err != nil {
		return err
	}
	if ping != "PONG" {
		return errors.New("connect failed")
	}

	rdb = db
	return nil
}


package dao

import (
	"github.com/go-redis/redis/v8"
)

var Rd0 *redis.Client

func GetRdbClient() *redis.Client {
	return Rd0
}

func InitRedis() {
	Rd0 = redis.NewClient(&redis.Options{
		Addr:     "b.y1ng.vip:46379",
		Password: "Byt3G0pheR51522zzwlwlbb",
		DB:       0, // 粉丝列表信息存入 DB0.
	})
}

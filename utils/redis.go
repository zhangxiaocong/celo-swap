package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	expireTime = 600
	rdb  *redis.Client
)

func main() {
	Startredis()
}

func Startredis() {
	rdb = redis.NewClient(&redis.Options{Addr: RedisIp + ":" + RedisPort, Password: "",})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis连接错误")
	}
}

func SaveValue(value string) (error, bool) {
	println("update celo price...")
	err := rdb.Set(context.Background(), "celo", value, time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return nil, true
	}
	return err, false
}

func GetValue(key string )string {
	value, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("设置key失败",err)
	}
	println("celo saved price",value)
	return value
}

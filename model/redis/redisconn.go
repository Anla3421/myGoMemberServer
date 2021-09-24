package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisConn *redis.Client

func init() {
	fmt.Println("Redis initial")
	CreateConn()
	GetDBInfo()
}

func CreateConn() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisConn = rdb
	fmt.Println("SuccessFully Connected to Redis")

}

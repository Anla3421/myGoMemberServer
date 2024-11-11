package redis

import (
	"encoding/json"
	"fmt"
	"server/domain/repository/model/dao"
	"server/domain/repository/model/dto"

	"github.com/go-redis/redis/v8"
)

func GetDBInfo() {
	dbdata := dao.QueryInfoAll()
	for account := range dbdata {
		value, _ := json.Marshal(dbdata[account])

		fmt.Println("欲寫入的DB資料 :"+account, dbdata[account])

		err := redisConn.Set(ctx, account, string(value), 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := redisConn.Get(ctx, account).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("寫入redis的資料 :"+account, val)
	}
}

func QueryInfo(account string) (bool, dto.Response) {
	value, err := redisConn.Get(ctx, account).Result()
	if err == redis.Nil {
		return false, dto.Response{}
	}
	if err != nil {
		panic(err)
	}
	var response dto.Response
	json.Unmarshal([]byte(value), &response)
	return true, response
}

func DeleteInfo(account string) {
	_, err := redisConn.Get(ctx, account).Result()
	if err == redis.Nil {
		fmt.Println("not found in redis")
		return
	}
	if err != nil {
		panic(err)
	}
	redisConn.Del(ctx, account)
}

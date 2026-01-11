package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"fmt"
	
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Redis 地址
	Password: "",               // 如果没密码就留空
})
const REDIS_PREFIX = "third-service:"
var SHORTER_DURATION = 1 * time.Second
var LONGER_DURATION = 6 * time.Second




var db = map[string]string{
	"1001": "LegoDeathStarPack",
	"1002": "LegoModularBuildings", // this one
	"1003": "JellyCatBunnyXL",      // and this one, are popular
	"1004": "JellyCatBunnyL",
	"1005": "LegoMillenniumFalcon",
}
var checkDB = 0


func ttl(key string) time.Duration {
	switch key {
	case "1002", "1003":
		return LONGER_DURATION
	default:
		return SHORTER_DURATION
	}
}



func main() {

	setTTL()



	simulateFetches(20, 200 * time.Millisecond)

	fmt.Println("Now let's change strategy to 'all man being equal'.")
	SHORTER_DURATION = 3 * time.Second
	LONGER_DURATION = 3 * time.Second

	simulateFetches(20, 200 * time.Millisecond)


	


	

}

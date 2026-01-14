package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	
	
	
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Redis 地址
	Password: "",               // 如果没密码就留空
})
const REDIS_PREFIX = "third-service:"



var db = map[string]string{
	"1001": "LegoDeathStarPack",
	"1002": "LegoModularBuildings", // this one 
	"1003": "JellyCatBunnyXL",      // and this one, are popular
	"1004": "JellyCatBunnyL",
	"1005": "LegoMillenniumFalcon",
}
var checkDB = 0









func main() {

	

	simulateHotColdVsEqualTTL()



	


	


	

}

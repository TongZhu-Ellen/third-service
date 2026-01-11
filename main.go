package main

import (
	"context"
    "fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"math/rand"
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // Redis 地址
	Password: "",               // 如果没密码就留空
})
const REDIS_PREFIX = "third-service:"
const SHORTER_DURATION = 10 * time.Second
const LONGER_DURATION = 60 * time.Second




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

	totalExperience := 50
    for i := 0; i < totalExperience; i++ {
        var key string
		hot := []string{"1002", "1003"}
		cold := []string{"1001", "1004", "1005"}
        if rand.Float64() < 0.8 { // 80%概率热门
            key = hot[rand.Intn(len(hot))]
        } else { // 20%概率冷门
            key = cold[rand.Intn(len(cold))]
        }

        val := GetProduct(key)
		fmt.Println(i+1, key, "=>", val)

        time.Sleep(500 * time.Millisecond)
	}
		

		fmt.Println("Amount", totalExperience, "total experience, we have check DB", checkDB, "times. ")
		fmt.Println("DB checking rate ==", float64(checkDB)/float64(totalExperience))




	

}

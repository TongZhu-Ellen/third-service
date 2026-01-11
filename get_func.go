package main



import (
	"github.com/redis/go-redis/v9"
	"log"

)


func GetProduct(key string) string {
	val, err := redisClient.Get(ctx, REDIS_PREFIX + key).Result() // to make it a go return

	if err == redis.Nil {
		val = db[key]
		checkDB++
		err = redisClient.Set(ctx, REDIS_PREFIX + key, val, 0).Err()
		if err != nil {
			log.Println("Redis set error:", err)
		}


    } else if err != nil {
		log.Println("Redis get error:", err)
		return ""
    }

	return val
}

package main



import (
	"github.com/redis/go-redis/v9"
	"log"

)


func GetProduct(key string) string {
	val, err := redisClient.Get(ctx, REDIS_PREFIX+key).Result() // to make it a go return

	if err == redis.Nil {

		lp := lockOf(key)
		lp.Lock()
		defer lp.Unlock()
		
		val, err = redisClient.Get(ctx, REDIS_PREFIX+key).Result()
		    if err == redis.Nil {
				val = db[key]
				checkDB++
			}
		

    } else if err != nil {
		log.Println("Redis get error:", err)
		return ""
    }

	

	err = redisClient.Set(ctx, REDIS_PREFIX+key, val, ttlOf(key)).Err()
	if err != nil {
		log.Println("Redis set error:", err)
	}

	return val
}


// this has no concurrent protection and hence "naive"!
func NaiveGetProduct(key string) string {
	
	val, err := redisClient.Get(ctx, REDIS_PREFIX+key).Result() // to make it a go return

	if err == redis.Nil {
		val = db[key]
		checkDB++
		


    } else if err != nil {
		log.Println("Redis get error:", err)
		return ""
    }

	err = redisClient.Set(ctx, REDIS_PREFIX+key, val, ttlOf(key)).Err()
	if err != nil {
		log.Println("Redis set error:", err)
	}

	return val
}


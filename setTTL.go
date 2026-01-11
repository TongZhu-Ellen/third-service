package main

import (

	
	"log"
)






func setTTL() {


	for key, val := range db {

		err := redisClient.Set(ctx, REDIS_PREFIX + key, val, ttl(key)).Err()
		if err != nil {
			log.Fatal(err)
		}
	}


	
}
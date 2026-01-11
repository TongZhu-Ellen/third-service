package main

import (
	"time"
	"fmt"
	"math/rand"
)

func simulateFetches(totalExperience int, sleepTime time.Duration) {


	

    for i := 0; i < totalExperience; i++ {
        var key string
		hot := []string{"1002", "1003"}
		cold := []string{"1001", "1004", "1005"}

        if rand.Float64() < 0.8 { // 80%概率热门
            key = hot[rand.Intn(len(hot))]
        } else { // 20%概率冷门
            key = cold[rand.Intn(len(cold))]
        }

        _ = GetProduct(key)
		// fmt.Println(i+1, key, "=>", val)

        time.Sleep(sleepTime)
	}
		

		fmt.Println("Amount", totalExperience, "total experience, we have check DB", checkDB, "times. ")
		fmt.Println("DB checking rate ==", float64(checkDB)/float64(totalExperience))



}
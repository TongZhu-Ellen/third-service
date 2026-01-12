package main

import (

	"fmt"
	"time"
	"math/rand"
	"log"
	"sync"


)

func fetch(totalExperience int, sleepTime time.Duration) {
	var key string
	hot := []string{"1002", "1003"}
	cold := []string{"1001", "1004", "1005"}

	for i := 1; i <= totalExperience; i++ {

        if rand.Float64() < 0.8 { // 80%概率热门
            key = hot[rand.Intn(len(hot))]
        } else { // 20%概率冷门
            key = cold[rand.Intn(len(cold))]
        }
        GetProduct(key)
        time.Sleep(sleepTime)
	}
		

		fmt.Println("Amount", totalExperience, "total experience, we have check DB", checkDB, "times. ")
		fmt.Println("DB checking rate ==", float64(checkDB)/float64(totalExperience))
}

func simulateHotColdVsEqualTTL() {


	totalExperience := 100
	sleepTime := 10 * time.Millisecond

	fetch(totalExperience, sleepTime)

	



	fmt.Println("==============")
	
	err := redisClient.FlushAll(ctx).Err()
	if err != nil {
		log.Fatal("Redis flush error:", err)
	} 
	
	SHORTER_DURATION = 30 * time.Millisecond
	LONGER_DURATION = 30 * time.Millisecond
	checkDB = 0

	fmt.Println("Now let's try again with equal TTL for everyone! ")



	fetch(totalExperience, sleepTime)


	


}



func simulateProtectedVsNaiveGet() {

	concurrency := 10

	wgp := &sync.WaitGroup{}

	wgp.Add(concurrency)

	for j := 1; j <= concurrency; j++ {

		go func() {
			defer wgp.Done()
			GetProduct("1001")
		}()
	}

	wgp.Wait()
	fmt.Println("Amount", concurrency, "total experience, we have check DB", checkDB, "times. ")
	fmt.Println("DB checking rate ==", float64(checkDB)/float64(concurrency))




	fmt.Println("=====================")
	err := redisClient.FlushAll(ctx).Err()
	if err != nil {
		log.Fatal("Redis flush error:", err)
	} 
	checkDB = 0
	fmt.Println("Now let's use the naive get function instead... ")

	wgp.Add(concurrency)

	for j := 1; j <= concurrency; j++ {

		go func() {
			defer wgp.Done()
			NaiveGetProduct("1001")
		}()
	}

	wgp.Wait()
	fmt.Println("Amount", concurrency, "total experience, we have check DB", checkDB, "times. ")
	fmt.Println("DB checking rate ==", float64(checkDB)/float64(concurrency))










}
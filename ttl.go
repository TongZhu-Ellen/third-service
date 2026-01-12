package main

import (

	
	
	"time"
)


var SHORTER_DURATION = 10 * time.Millisecond
var LONGER_DURATION = 60 * time.Millisecond
// there two are suppose to be const, but... for the sake of testing I make it var!


func ttlOf(key string) time.Duration {
	switch key {
	case "1002", "1003":
		return LONGER_DURATION
	default:
		return SHORTER_DURATION
	}
}






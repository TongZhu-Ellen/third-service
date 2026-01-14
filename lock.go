package main

import (
	"sync"
)

var individualLPs sync.Map // key -> *sync.Mutex

func lockOf(key string) *sync.Mutex {
	// LoadOrStore 返回两个值：已存在的值 or 新创建的值，以及是否已存在
	lp, _ := individualLPs.LoadOrStore(key, &sync.Mutex{})
	return lp.(*sync.Mutex)
}

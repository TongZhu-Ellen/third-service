package main 

import (
	"sync"
)

var individualLPs = make(map[string]*sync.Mutex) // lock pointers
var grobalLP = &sync.Mutex{}

func lockOf(key string) *sync.Mutex {

	grobalLP.Lock()
	defer grobalLP.Unlock()

	_, exists := individualLPs[key]
	if !exists {
		individualLPs[key] = &sync.Mutex{}
	}
	
	return individualLPs[key]
}






package main

import (
	"fmt"
	"sync"
	"time"
)

var comData int = 0

func main() {
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				lock.Lock()
				comData++
				lock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(comData)
}

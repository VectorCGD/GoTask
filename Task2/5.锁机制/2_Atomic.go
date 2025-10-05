package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var comData atomic.Int32

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				comData.Add(1)
			}
		}()
	}
	time.Sleep(time.Second * 3)

	fmt.Println(comData.Load())
}

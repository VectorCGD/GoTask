package main

import (
	"fmt"
	"math/rand"
	"time"
)

func task(target int, c chan<- float64) {
	st := time.Now()
	count := 0
	for i := 0; i < target; i++ {
		count++
	}
	dt := time.Since(st)

	c <- dt.Seconds()
}

func main() {
	times := make(chan float64, 10)
	for i := 0; i < 10; i++ {
		workload := rand.Intn(1000000-100000+1) + 100000
		go task(workload, times)
	}

	time.Sleep(time.Second * 5)
	close(times)

	for k := range times {
		fmt.Println(k)
	}
}

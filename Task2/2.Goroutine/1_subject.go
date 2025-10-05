package main

import (
	"fmt"
	"time"
)

func oddEvenNumbers(condition func(int) bool) {
	for i := 0; i < 10; i++ {
		if condition(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	go oddEvenNumbers(func(i int) bool { return i%2 == 0 })
	go oddEvenNumbers(func(i int) bool { return i%2 == 1 })
	time.Sleep(time.Second)
}

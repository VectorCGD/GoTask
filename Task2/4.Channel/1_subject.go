package main

import (
	"fmt"
)

func main() {
	completed := false
	ch := make(chan int)
	//input
	go func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}(ch)
	//output
	go func(c chan int) {
		for v := range c {
			fmt.Println(v)
		}
		completed = true
	}(ch)

	for !completed {
	}
}

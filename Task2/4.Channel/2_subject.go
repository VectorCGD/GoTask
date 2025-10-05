package main

import (
	"fmt"
)

var completed bool = false

func producer(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func consumer(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	completed = true
}
func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)

	for !completed {
	}
}

package main

import "fmt"

func add(numptr *int) {
	*numptr += 10
}

func main() {
	n := 5
	add(&n)
	fmt.Println(n)
}

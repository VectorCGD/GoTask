package main

import "fmt"

func doubleNumbers(nums *[]int) {
	for i, _ := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	slice := []int{2, 3, 5, 8, 9}
	doubleNumbers(&slice)
	fmt.Println(slice)
}

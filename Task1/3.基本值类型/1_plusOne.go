package main

import "fmt"

func plusOne(digits []int) []int {
	p := -1
	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			p = i
			break
		}
	}
	if p == -1 {
		result := make([]int, 1, len(digits)+1)
		result[0] = 1
		return append(result, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{8, 9, 8, 9}))
}

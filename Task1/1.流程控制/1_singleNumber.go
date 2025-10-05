package main

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, exist := m[nums[i]]; exist {
			m[nums[i]] = false
		} else {
			m[nums[i]] = true
		}
	}
	var result int
	for k, v := range m {
		if v {
			result = k
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 3, 1, 7, 14, 6, 5, 7, 6, 10, 14}
	fmt.Println(singleNumber(nums))
}

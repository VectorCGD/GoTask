package main

import "fmt"

func removeDuplicates(nums []int) (int, []int) {
	result := make([]int, 0, len(nums))
	result = append(result, nums[0])
	for _, v := range nums {
		if result[len(result)-1] == v {
			continue
		} else {
			result = append(result, v)
		}
	}
	return len(result), result
}

func main() {
	c, _ := removeDuplicates([]int{2, 2, 3, 3, 3, 4, 4, 5, 5, 7, 8, 8, 9})
	fmt.Println(c)
}

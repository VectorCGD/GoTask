package main

import "fmt"

func twoSum(nums []int, target int) []int {
	resultMap := make(map[int][]int, len(nums))
	for index, v := range nums {
		if _, exist := resultMap[v]; !exist {
			resultMap[v] = make([]int, 0, 2)
			resultMap[v] = append(resultMap[v], index)
			comNumber := target - v
			if r, exist := resultMap[comNumber]; exist {
				resultMap[v] = append(resultMap[v], r[0])
				return resultMap[v]
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 5, 3, 4, 8, 7, 9}, 13))
}

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])
	for interval, i := result[0], 1; i < len(intervals); i++ {
		if interval[1] >= intervals[i][0] {
			if interval[1] < intervals[i][1] {
				interval[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
			interval = result[len(result)-1]
		}
	}
	return result
}

func main() {
	fmt.Println(merge([][]int{{4, 9}, {2, 5}, {3, 6}, {10, 11}}))

}

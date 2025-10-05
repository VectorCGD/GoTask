package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var strBuild strings.Builder
comFor:
	for i, next := 0, true; next; i++ {
		var char byte
		for n := 0; n < len(strs); n++ {
			if i >= len(strs[n]) {
				next = false
				continue comFor
			}
			if char == 0 {
				char = strs[n][i]
			}
			if char != strs[n][i] {
				next = false
				continue comFor
			}

		}
		strBuild.WriteString(string(char))
	}
	return strBuild.String()
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"adbdq", "adb", "adbdcsw"}))
}

package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for strLength := len(s) + 1; len(s) > 0 && len(s) < strLength; {
		strLength = len(s)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
	}
	return len(s) == 0
}

func main() {
	fmt.Println(isValid("{()}({})[{}][({})]"))
}

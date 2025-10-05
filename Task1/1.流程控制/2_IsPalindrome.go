package main

import "fmt"

func IsPalindrome(x int) bool {
	reverse := 0
	for i := x; i > 0; i /= 10 {
		reverse = reverse*10 + i%10
	}
	return reverse == x
}

func main() {
	fmt.Println(IsPalindrome(12345321))
}

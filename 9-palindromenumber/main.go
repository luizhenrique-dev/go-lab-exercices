package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := -53521
	input := 123321
	p := isPalindrome(input)
	conjunction := map[bool]string{true: "is", false: "isn't"}[p]
	fmt.Printf("The number %d %s a palindrome", input, conjunction)
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	xStr := strconv.Itoa(x)
	xLength := len(xStr)
	var sb strings.Builder
	for i := xLength; i > 0; i-- {
		sb.WriteString(xStr[i-1 : i])
	}
	p, _ := strconv.Atoi(sb.String())
	return p == x
}

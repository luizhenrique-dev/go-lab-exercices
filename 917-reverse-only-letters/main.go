package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "Test1ng-Leet=code-Q!"
	r := reverseOnlyLetters(str)
	fmt.Println(r)
}

func reverseOnlyLetters(str string) string {
	idxLeft, idxRight := 0, len(str)-1
	lettersRegex := regexp.MustCompile(`^[A-Z|a-z]$`)
	strChars := []rune(str)
	for idxLeft < idxRight {
		if !lettersRegex.MatchString(string(strChars[idxLeft])) {
			idxLeft++
		} else if !lettersRegex.MatchString(string(strChars[idxRight])) {
			idxRight--
		} else {
			strChars[idxLeft], strChars[idxRight] = strChars[idxRight], strChars[idxLeft]
			idxLeft++
			idxRight--
		}

	}
	return string(strChars)
}

func reverseOnlyLetters2(str string) string {
	lettersRegex := regexp.MustCompile(`^[A-Z|a-z]$`)
	var onlyLettersRevBuilder strings.Builder
	for i := len(str); i > 0; i-- {
		charStr := str[i-1 : i]
		if lettersRegex.MatchString(charStr) {
			onlyLettersRevBuilder.WriteString(charStr)
		}
	}

	var rb strings.Builder
	onlyLettersRev := onlyLettersRevBuilder.String()
	lastIdx := 0
	for i := 0; i < len(str); i++ {
		charStr := str[i : i+1]
		if !lettersRegex.MatchString(charStr) {
			rb.WriteString(charStr)
		} else {
			rb.WriteString(string(onlyLettersRev[lastIdx]))
			lastIdx++
		}
	}
	return rb.String()
}

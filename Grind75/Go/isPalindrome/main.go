package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	pattern := regexp.MustCompile(`[a-zA-Z0-9]+`)
	s = strings.ToLower(strings.Join(pattern.FindAllString(s, -1), ""))

	if s == reverse(s) {
		return true
	}

	return false
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	test := "A man, a plan, a canal: Panama"
	if isPalindrome(test) {
		fmt.Println("passed")
	} else {
		fmt.Println("failed")
	}
}

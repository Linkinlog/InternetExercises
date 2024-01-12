package main

import (
	"fmt"
	"strings"
)

func isAnagram(s string, t string) bool {
	anagram := true
	if len(s) != len(t) {
		return false
	}
	for _,c := range t {
		if !strings.ContainsRune(s, c) {
			anagram = false
		} else {
			s = strings.Replace(s, string(c), "", 1)
		}
	}
	return anagram
}

func main() {
	string := "aacc"
	expected := "ccac"

	if isAnagram(string, expected) {
		fmt.Println("Fail")
	} else {
		fmt.Println("Success")
	}
}

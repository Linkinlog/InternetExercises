// palindrome provides utilities to determine if a string is a palindrome
package palindrome

import (
	"strings"
)

// IsPalindrome takes in a string and returns a boolean value indicating wether or not the given string is a palindrome
func IsPalindrome(str string) bool {
	// Set the string to lowercase for consistency
	lowercase := strings.ToLower(str)
	// Iterate over half of the string
	for i := 0; i < len(lowercase)/2; i++ {
		// if char isnt the same as its opposing char then return false
		if lowercase[len(lowercase)-1-i] != lowercase[i] {
			return false
		}
	}
	return true
}

// Package hashTagGen provides utilities relative to strings and hashtags
package hashtaggen

import (
	"errors"
	"unicode"
)

// HashTagGen will take a string and return either the hashtag variant of the string or an error
func HashTagGen(s string) (error, string) {
	if s == "" {
		return errors.New("Invalid string"), ""
	}
	var newStr []rune = []rune{'#'}
	shouldCapitalize := false
	// Iterate over the string
	for _, e := range s {
		if e != ' ' {
			if shouldCapitalize {
				// After space, ensure the next char is capitalized
				newStr = append(newStr, unicode.ToUpper(e))
				shouldCapitalize = false
			} else {
				// If not a space, append to the rune array
				newStr = append(newStr, e)
			}
		}
		// If space then skip
		if e == ' ' {
			shouldCapitalize = true
		}

	}

	return nil, string(newStr)
}

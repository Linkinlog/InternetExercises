// Package validParens has utilities to validate parenthesis
package validParens

// Validate validates if the string has valid parenthesis or not
func Validate(parens string) bool {
	// Running total of '(' without ')'
	var openParens int
	for _, e := range parens {
		// Loop over the string, increment if '(', decrement if ')'
		if e == 41 {
			openParens--
		} else if e == 40 {
			openParens++
		}
		// openParens will be negative if a ')' is hit before '('
		if openParens < 0 {
			return false
		}
	}
	return openParens == 0
}

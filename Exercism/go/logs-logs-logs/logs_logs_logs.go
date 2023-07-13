package logs

import (
	"strings"
)

// The Application identifies the application emitting the given log.
func Application(log string) string {
	str := "default"
	for _, el := range log {
		switch el {
		case '❗':
			str = "recommendation"
			return str
		case '🔍':
			str = "search"
			return str
		case '☀':
			str = "weather"
			return str
		}
	}
	return str
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether the number of characters in the log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if len([]rune(log)) > limit {
		return false
	}
	return true
}

package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[TRC]|^\[ERR]|^\[DBG]|^\[INF]|^\[WRN]|^\[FTL]`).Match([]byte(text))
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	return len(regexp.MustCompile(`(?i)"[^"]*password[^"]*"`).FindAllString(strings.Join(lines, " "), -1))
}

func RemoveEndOfLineText(text string) string {
	return string(regexp.MustCompile(`end-of-line\d+`).ReplaceAll([]byte(text), []byte("")))
}

func TagWithUserName(lines []string) []string {
	reg := regexp.MustCompile(`User\s+([^ ]*)\s`)

	for i, el := range lines {
		matches := reg.FindStringSubmatch(el)

		if len(matches) > 1 {
			name := matches[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", name, lines[i])
		}
	}
	return lines
}

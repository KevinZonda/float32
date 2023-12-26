package utils

import (
	"strings"
	"unicode"
)

func CleanStr(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		if r == '\n' || r == '\r' || r == '\t' {
			return r
		}
		return -1
	}, s)
}

func StrMaxLen(str string, maxLen int) string {
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen])
}

func StrMaxLenSmart(str string, maxLen int) string {

	bl := float64(len([]byte(str)))
	rl := float64(len([]rune(str)))
	// if is latin, then double len
	needDouble := (1.0*(bl-rl))/bl < 0.2
	if needDouble {
		maxLen = int(2.5 * float64(maxLen))
	}
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen])
}

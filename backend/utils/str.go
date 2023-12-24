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

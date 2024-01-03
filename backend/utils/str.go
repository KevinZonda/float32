package utils

import (
	"io"
	"strings"
	"unicode"
)

func HasTails(str string, tails ...string) bool {
	for _, tail := range tails {
		if strings.HasSuffix(str, tail) {
			return true
		}
	}
	return false
}

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

func StrMaxLenSmart(str string, maxLen int, tail string) string {
	// if is latin, then double len
	diffRate := StrByteRuneDiffRate(str)
	if diffRate < 0.2 {
		maxLen = int(4 * float64(maxLen))
	} else if diffRate < 0.4 {
		maxLen = int(2 * float64(maxLen))
	}
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen]) + tail
}

func StrByteRuneDiffRate(str string) float64 {
	bl := float64(len([]byte(str)))
	rl := float64(len([]rune(str)))
	// if is latin, then double len
	return (1.0 * (bl - rl)) / bl
}

func StrContains(c string, ss ...string) bool {
	if ss == nil {
		return false
	}

	for _, s := range ss {
		if s == c {
			return true
		}
	}
	return false
}

func IndexOfRunes(rs []rune, r rune) int {
	for idx, r1 := range rs {
		if r == r1 {
			return idx
		}
	}
	return -1
}

func HasSensitiveWords(str string) bool {
	for _, s := range sensitiveWords {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

var sensitiveWords = []string{
	"微信：",
	"微信:",
	"微信电话：",
	"微信电话:",
	"同城",
	"少妇",
	"Google My Maps",
}

func writeSplitByRune(w io.Writer, sb, buf *strings.Builder, delta string, subStr rune) (needContinue bool) {
	rs := []rune(delta)
	idx := IndexOfRunes(rs, subStr)
	if idx < 0 {
		return
	}
	toPrint := buf.String() + string(rs[:idx+1])
	w.Write([]byte(toPrint))
	buf.Reset()
	sb.WriteString(string(rs[idx+1:]))
	needContinue = true
	return
}

func WriteSplitByRune(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if writeSplitByRune(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

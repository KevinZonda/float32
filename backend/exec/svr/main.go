package main

import (
	"github.com/KevinZonda/float32/utils"
	"io"
	"strings"
)

func main() {
	initAll()
	g.POST("/query", queryQuestion)
	g.GET("/history", ginHistory)
	startGin()
}

func writeBySubStr(w io.Writer, sb, buf *strings.Builder, delta string, subStr rune) (needContinue bool) {
	rs := []rune(delta)
	idx := utils.IndexOfRunes(rs, subStr)
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

func writeBySubStrs(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if writeBySubStr(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

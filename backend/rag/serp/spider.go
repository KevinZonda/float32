package serp

import (
	"github.com/KevinZonda/float32/utils"
	"strings"
)

type Spider interface {
	Search(urls ...string) (resp []SpiderResult)
}

type SpiderResult struct {
	Title       string
	Url         string
	Description string
	Content     string
	Error       error
}

func (r SpiderResult) String(maxLen int) string {
	sb := strings.Builder{}
	sb.WriteString("* URL: ")
	sb.WriteString(r.Url)
	sb.WriteString("    ")
	sb.WriteString("Title: ")
	sb.WriteString(r.Title)
	sb.WriteString("\n")
	//first 1000 chars
	sb.WriteString(utils.StrMaxLenSmart(r.Content, maxLen, ""))
	sb.WriteString("\n")
	return sb.String()
}

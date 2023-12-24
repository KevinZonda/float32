package rag

import (
	"github.com/KevinZonda/float32/rag/serp"
	"log"
	"os"
	"strings"
)

const SearchPerItemMaxLen = 1000
const SearchMaxItemCount = 5

func Search(query string) string {
	gs := serp.NewGoogleSearch(os.Getenv("SERP_DEV"))
	resp, err := gs.Search(query)
	if err != nil {
		return err.Error()
	}
	var urls []string
	for _, r := range resp.Result {
		urls = append(urls, r.Link)
	}
	if len(urls) == 0 {
		return ""
	}
	headOfSearch := arrMaxLen[string](urls, SearchMaxItemCount)
	spider := serp.NewSimpleSpider()
	results := spider.Search(headOfSearch...)
	sb := strings.Builder{}
	for _, r := range results {
		if r.Error != nil {
			log.Println("FAILED", r.Url, r.Error)
			continue
		}
		sb.WriteString("* URL: ")
		sb.WriteString(r.Url)
		sb.WriteString("    ")
		sb.WriteString("Title: ")
		sb.WriteString(r.Title)
		sb.WriteString("\n")
		//first 1000 chars
		sb.WriteString(strMaxLen(r.Content, SearchPerItemMaxLen))
		sb.WriteString("\n\n")
	}
	return sb.String()
}

func strMaxLen(str string, maxLen int) string {
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen])
}

func arrMaxLen[T any](str []T, maxLen int) []T {
	if len(str) <= maxLen {
		return str
	}
	return str[:maxLen]
}

func MapProgLang(progLang string) string {
	switch progLang {
	case "Go":
		return "Golang"
	default:
		return progLang
	}
}

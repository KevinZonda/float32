package main

import (
	"encoding/json"
	"github.com/KevinZonda/float32/rag/serp"
	"log"
	"os"
	"strings"
)

func search(query string) string {
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
	first5 := arrMaxLen[string](urls, 5)
	spider := serp.NewSimpleSpider()
	results := spider.Search(first5...)
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
		sb.WriteString(strMaxLen(r.Content, 1000))
		sb.WriteString("\n\n")
	}
	return sb.String()
}

func jout(v any) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	println(string(bs))
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

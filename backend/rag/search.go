package rag

import (
	"github.com/KevinZonda/float32/rag/serp"
	"log"
	"os"
	"strings"
	"time"
)

const SearchPerItemMaxLen = 500
const SearchMaxItemCount = 5

func hasTails(str string, tails ...string) bool {
	for _, tail := range tails {
		if strings.HasSuffix(str, tail) {
			return true
		}
	}
	return false
}

func Search(query string) string {
	beforeGoogleTime := time.Now()
	gs := serp.NewGoogleSearch(os.Getenv("SERP_DEV"))
	resp, err := gs.Search(query)
	if err != nil {
		return err.Error()
	}
	var urls []string
	for _, r := range resp.Result {
		if len(urls) >= SearchMaxItemCount {
			break
		}
		link := r.Link
		if hasTails(link, ".pdf", ".doc", ".docx", ".ppt", ".pptx", ".xls") {
			continue
		}
		if strings.Contains(link, "gov.cn") {
			continue
		}
		urls = append(urls, r.Link)

	}
	if len(urls) == 0 {
		return ""
	}

	afterGoogleTime := time.Now()
	spider := serp.NewSimpleSpider()
	results := spider.Search(urls...)
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
	spiderTime := time.Now()
	log.Println("Search Time:", afterGoogleTime.Sub(beforeGoogleTime), "Spider Time:", spiderTime.Sub(afterGoogleTime))
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

package rag

import (
	"github.com/KevinZonda/float32/rag/serp"
	"github.com/KevinZonda/float32/utils"
	"log"
	"os"
	"strings"
	"time"
)

const SearchPerItemMaxLen = 500
const SearchMaxItemCount = 8
const SearchMaxIncludeInContext = 5

func hasTails(str string, tails ...string) bool {
	for _, tail := range tails {
		if strings.HasSuffix(str, tail) {
			return true
		}
	}
	return false
}

type urlInfo struct {
	Title       string
	Description string
}

func SearchRaw(country, query string) ([]serp.SpiderResult, error) {
	beforeGoogleTime := time.Now()
	gs := serp.NewGoogleSearch(os.Getenv("SERP_DEV"))
	resp, err := gs.Search(country, query)
	if err != nil {
		return nil, err
	}
	var urls []string
	urlMap := make(map[string]urlInfo)
	for _, r := range resp.Result {
		if len(urls) >= SearchMaxItemCount {
			break
		}
		if strings.HasPrefix(r.Title, "[PDF]") {
			continue
		}
		link := r.Link
		if hasTails(link, ".pdf", ".doc", ".docx", ".ppt", ".pptx", ".xls") {
			continue
		}
		if strings.Contains(link, "gov.cn") {
			continue
		}
		urls = append(urls, link)
		urlMap[link] = urlInfo{
			Title:       r.Title,
			Description: r.Snippet,
		}
	}
	if len(urls) == 0 {
		return nil, nil
	}

	afterGoogleTime := time.Now()

	spider := serp.NewSimpleSpider()
	results := spider.Search(urls...)
	spiderTime := time.Now()
	for i, r := range results {
		if r.Error != nil {
			continue
		}
		info := urlMap[r.Url]
		results[i].Title = info.Title
		results[i].Description = info.Description
	}

	if len(results) > SearchMaxIncludeInContext {
		results = results[:SearchMaxIncludeInContext]
	}

	log.Println("Search Time:", afterGoogleTime.Sub(beforeGoogleTime), "Spider Time:", spiderTime.Sub(afterGoogleTime))

	return results, nil
}

func Search(query string) string {
	results, err := SearchRaw("us", query)
	if err != nil {
		return ""
	}
	return SearchResultsToText(results)
}

func SearchResultsToText(results []serp.SpiderResult) string {
	sb := strings.Builder{}
	if len(results) == 0 {
		return ""
	}
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
		sb.WriteString(utils.StrMaxLenSmart(r.Content, SearchPerItemMaxLen, ""))
		sb.WriteString("\n\n")
	}
	return sb.String()
}

func MapProgLang(progLang string) string {
	switch progLang {
	case "Go":
		return "Golang"
	default:
		return progLang
	}
}

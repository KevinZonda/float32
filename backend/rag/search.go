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

type urlInfo struct {
	Title       string
	Description string
}

func canBeAddTo(r serp.GoogleSearchResponseItem) bool {
	if strings.HasPrefix(r.Title, "[PDF]") {
		return false
	}
	link := r.Link
	if utils.HasTails(link, ".pdf", ".doc", ".docx", ".ppt", ".pptx", ".xls") {
		return false
	}
	if strings.Contains(link, "gov.cn") {
		return false
	}
	if utils.HasSensitiveWords(r.Title) {
		return false
	}
	//if utils.HasSensitiveWords(r.Snippet) {
	//	return false
	//}
	return true
}

type SearchRawItem struct {
	SpiderResults []serp.SpiderResult
	Related       []string
}

func SearchRaw(country, locale, query string) (SearchRawItem, error) {
	rst := SearchRawItem{}
	beforeGoogleTime := time.Now()
	gs := serp.NewGoogleSearch(os.Getenv("SERP_DEV"))
	resp, err := gs.Search(country, locale, query)
	if err != nil {
		return SearchRawItem{}, err
	}
	rst.Related = resp.RelatedSearchStrs()
	var urls []string
	urlMap := make(map[string]urlInfo)
	for _, r := range resp.Result {
		if len(urls) >= SearchMaxItemCount {
			break
		}
		if !canBeAddTo(r) {
			continue
		}
		urls = append(urls, r.Link)
		urlMap[r.Link] = urlInfo{
			Title:       r.Title,
			Description: r.Snippet,
		}
	}
	if len(urls) == 0 {
		return rst, nil
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
	rst.SpiderResults = results

	log.Println("Search Time:", afterGoogleTime.Sub(beforeGoogleTime), "Spider Time:", spiderTime.Sub(afterGoogleTime))

	return rst, nil
}

func Search(query string) string {
	results, err := SearchRaw("us", "", query)
	if err != nil {
		return ""
	}
	return SearchResultsToText(results.SpiderResults)
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
		sb.WriteString(r.String(SearchPerItemMaxLen))
		sb.WriteString("\n")
	}
	return sb.String()
}

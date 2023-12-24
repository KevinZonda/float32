package serp

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
	"sync"
)

type SimpleSpider struct {
	hc *http.Client
}

func (s *SimpleSpider) Search(urls ...string) (results []SpiderResult) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	syncMap := sync.Map{}
	for _idx, _url := range urls {
		go func(idx int, url string) {
			defer wg.Done()
			result := SpiderResult{
				Url: url,
			}
			resp, err := s.get(url)
			if err != nil {
				result.Error = err
				goto store
			}
			if resp.StatusCode != http.StatusOK {
				result.Error = errors.New("status code not 200")
				goto store
			}
			defer resp.Body.Close()
			result = _scrab(resp.Body)
			result.Url = url
		store:
			syncMap.Store(idx, result)
		}(_idx, _url)
	}
	wg.Wait()
	syncMap.Range(func(key, value interface{}) bool {
		results = append(results, value.(SpiderResult))
		return true
	})
	return
}

func NewSimpleSpider() Spider {
	return &SimpleSpider{
		hc: &http.Client{},
	}
}

func (s *SimpleSpider) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	if err != nil {
		return
	}
	return s.hc.Do(req)
}

func _scrab(r io.Reader) (resp SpiderResult) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		resp.Error = err
		return
	}
	doc.Find("title").Each(func(i int, selection *goquery.Selection) {
		resp.Title = selection.Text()
	})
	doc.Find("meta[name=description]").Each(func(i int, selection *goquery.Selection) {
		resp.Description, _ = selection.Attr("content")
	})
	doc.Find("body").Each(func(i int, selection *goquery.Selection) {
		// ignore footer
		selection.Find("footer").Remove()
		selection.Find("header").Remove()
		selection.Find("noscript").Remove()
		// ignore script
		selection.Find("script").Remove()
		// ignore style
		selection.Find("style").Remove()
		resp.Content = selection.Text()
		resp.Content = strings.TrimSpace(resp.Content)
	})
	return
}

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
			result = _scrab(_url, resp.Body)
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

func _scrab(url string, r io.Reader) (resp SpiderResult) {
	resp.Url = url
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
		defaultRemover(selection)
		if strings.Contains(url, "geeksforgeeks.org") {
			geeksforgeeks(selection)
		}
		if strings.Contains(url, "stackoverflow.com") {
			stackOveflow(selection)
		}
		resp.Content = selection.Text()
		resp.Content = strings.TrimSpace(resp.Content)
	})
	return
}

func defaultRemover(selection *goquery.Selection) {
	selection.Find("footer").Remove()
	selection.Find("header").Remove()
	selection.Find("noscript").Remove()
	selection.Find("script").Remove()
	selection.Find("[class*=footer]").Remove()
	selection.Find("[class*=sidebar]").Remove()
	selection.Find("[class*=search]").Remove()
	selection.Find("[class*=nav]").Remove()
	selection.Find("a[href*=twitter.com]").Remove()
	selection.Find("a[href*=facebook.com]").Remove()
	selection.Find("a[href*=linkedin.com]").Remove()
	selection.Find("[id*=footer]").Remove()
	selection.Find("[id*=sidebar]").Remove()
	selection.Find("nav").Remove()
	selection.Find("style").Remove()
}

func geeksforgeeks(selection *goquery.Selection) {
	defaultRemover(selection)
	selection.Find("[class*=footer]").Remove()
	selection.Find("[class*=header-main__container]").Remove()
	selection.Find("[class*=header-main__slider]").Remove()
	selection.Find("[class*=header-sidebar__wrapper]").Remove()
	selection.Find("[class*=gfg-footer]").Remove()
	selection.Find("[class*=article--recommended]").Remove()
	selection.Find("[class*=cookie-consent]").Remove()
}

func stackOveflow(selection *goquery.Selection) {
	defaultRemover(selection)
	//remove id
	selection.Find("[id*=left-sidebar]").Remove()
	selection.Find("[class*=js-dismissable-hero]").Remove()
	selection.Find("[id*=post-form]").Remove()

}

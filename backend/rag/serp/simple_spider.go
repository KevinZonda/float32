package serp

import (
	"bytes"
	"errors"
	"github.com/go-shiori/go-readability"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type SimpleSpider struct {
	hc      *http.Client
	Timeout time.Duration
	HttpReq func(req *http.Request)
}

func (s *SimpleSpider) Search(urls ...string) (results []SpiderResult) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	syncMap := sync.Map{}
	for _idx, _url := range urls {
		go func(idx int, url string) {
			log.Println("Spider", url, "time", time.Now())
			var body []byte
			spiderStart := time.Now()
			spiderEnd := time.Now()
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
			body, err = readAllWithTimeout(resp.Body, s.Timeout)
			if err != nil {
				result.Error = err
				spiderEnd = time.Now()
				goto store
			}
			result = readabilityScrab(url, bytes.NewReader(body))
			spiderEnd = time.Now()
		store:
			syncMap.Store(idx, result)
			mutexTime := time.Now()
			log.Println("Spider Time:", spiderEnd.Sub(spiderStart), "Mutex Time:", mutexTime.Sub(spiderEnd))
		}(_idx, _url)
	}
	wg.Wait()
	syncMap.Range(func(key, value interface{}) bool {
		results = append(results, value.(SpiderResult))
		return true
	})
	return
}

func readAllWithTimeout(r io.Reader, timeout time.Duration) (result []byte, err error) {
	ch := make(chan []byte)
	start := time.Now()
	go func() {
		result, err = io.ReadAll(r)
		ch <- result
	}()
	select {
	case <-time.After(timeout):
		err = errors.New("timeout" + time.Since(start).String())
	case result = <-ch:
	}
	return
}

func NewSimpleSpider() *SimpleSpider {
	return &SimpleSpider{
		hc: &http.Client{
			Timeout: 5 * time.Second,
		},
		Timeout: 5 * time.Second,
		HttpReq: GoogleBotHeader,
	}
}

func (s *SimpleSpider) SetTimeout(t time.Duration) {
	s.hc.Timeout = t
}

func (s *SimpleSpider) GetTimeOut() time.Duration {
	return s.hc.Timeout
}

func (s *SimpleSpider) get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if s.HttpReq != nil {
		s.HttpReq(req)
	}
	if err != nil {
		return
	}
	return s.hc.Do(req)
}

func readabilityScrab(urlS string, r io.Reader) (resp SpiderResult) {
	resp.Url = urlS
	parsedUrl, _ := url.Parse(urlS)

	rd, err := readability.FromReader(r, parsedUrl)
	if err != nil {
		resp.Error = err
		return
	}
	resp.Title = strings.TrimSpace(rd.Title)
	resp.Description = strings.TrimSpace(rd.Excerpt)
	resp.Content = strings.Replace(strings.TrimSpace(rd.TextContent), "\n\n", "\n", -1)
	return
}

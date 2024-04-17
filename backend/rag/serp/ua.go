package serp

import (
	"math/rand/v2"
	"net/http"
)

const UA_GoogleBot = "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/120.0.0.0 Safari/537.36"
const UA_BingBot = "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Chrome//120.0.0.0 Safari/537.36"
const UA_BaiduBot = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html）"

func GoogleBotHeader(req *http.Request) {
	req.Header.Add("User-Agent", UA_GoogleBot)
	req.Header.Add("From", "googlebot(at)googlebot.com")
	req.Header.Add("Accept", "text/plain,text/html,*/*")
	req.Header.Add("Connection", "keep-alive")
}

func BingBotHeader(req *http.Request) {
	req.Header.Add("User-Agent", UA_BingBot)
	req.Header.Add("Accept", "text/plain,text/html,*/*")
	req.Header.Add("Connection", "keep-alive")
}

func BaiduBotHeader(req *http.Request) {
	req.Header.Add("User-Agent", UA_BaiduBot)
	req.Header.Add("Accept", "text/plain,text/html,*/*")
	req.Header.Add("Connection", "keep-alive")
}

func RandomBotHeader(req *http.Request) {
	n := rand.IntN(3)
	switch n {
	case 0:
		GoogleBotHeader(req)
	case 1:
		BingBotHeader(req)
	case 2:
		BaiduBotHeader(req)
	default:
		GoogleBotHeader(req)
	}
}

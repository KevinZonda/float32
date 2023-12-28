package serp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GoogleSearch struct {
	apiKey string
	hc     *http.Client
}

func NewGoogleSearch(apiKey string) *GoogleSearch {
	return &GoogleSearch{
		apiKey: apiKey,
		hc:     &http.Client{},
	}
}

type GoogleSearchResponse struct {
	//AnswerBox struct {
	//	Snippet string `json:"snippet"`
	//	Title   string `json:"title"`
	//	Link    string `json:"link"`
	//	Date    string `json:"date"`
	//} `json:"answerBox"`
	Result []GoogleSearchResponseItem `json:"organic"`
}

type GoogleSearchResponseItem struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Snippet   string `json:"snippet"`
	Position  int    `json:"position"`
	SiteLinks []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"sitelinks,omitempty"`
}

type _googleQuery struct {
	Query       string `json:"q"`
	Country     string `json:"gl,omitempty"`
	Locale      string `json:"hl,omitempty"`
	AutoCorrect bool   `json:"autocorrect,omitempty"`
	Page        int    `json:"page,omitempty"`
	NumOfResult int    `json:"num,omitempty"` // default 10, max 100
}

const serperUrl = "https://google.serper.dev/search"

func (s *GoogleSearch) Search(country, query string) (resp GoogleSearchResponse, err error) {
	method := "POST"

	reqM := _googleQuery{
		Query:   query,
		Country: country,
		//Locale:  "en-us",
	}
	reqB, err := json.Marshal(reqM)
	if err != nil {
		return
	}

	payload := bytes.NewReader(reqB)

	req, err := http.NewRequest(method, serperUrl, payload)

	if err != nil {
		return
	}
	req.Header.Add("X-API-KEY", s.apiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := s.hc.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resp)
	return
}

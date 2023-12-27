package main

import (
	"encoding/json"
	"github.com/KevinZonda/float32/rag/serp"
)

type MetaModel struct {
	Evidences []SearchItem `json:"evidences"`
}

type SearchItem struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (m MetaModel) Json() string {
	b, _ := json.Marshal(m)
	return string(b)
}

type Query struct {
	Question string `json:"question"`
	ProgLang string `json:"prog_lang"`
	Language string `json:"language"`
	Field    string `json:"field"`
}

func (q Query) Regularize() Query {
	switch q.Language {
	case "简体中文", "zh":
		q.Language = "简体中文"
	default:
		q.Language = "English"
	}
	switch q.Field {
	case "code", "med":
	default:
		q.Field = "code"
	}

	switch q.ProgLang {
	case "Go", "go":
		q.ProgLang = "golang"

	}
	return q
}

func newMeta(searched []serp.SpiderResult) MetaModel {
	var evi []SearchItem
	for _, r := range searched {
		if r.Error != nil {
			continue
		}
		evi = append(evi, SearchItem{
			Url:         r.Url,
			Title:       r.Title,
			Description: r.Description, // utils.StrMaxLenSmart(r.Description, 50, "..."),
		})
	}
	return MetaModel{
		Evidences: evi,
	}
}

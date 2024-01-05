package reqmodel

import (
	"github.com/KevinZonda/float32/utils"
	"strings"
)

type Query struct {
	Question string `json:"question"`
	ProgLang string `json:"prog_lang"`
	Language string `json:"language"`
	Field    string `json:"field"`
}

func (q Query) Regularize() Query {
	switch q.Language {
	case "简体中文", "zh":
		q.Language = "Chinese (Mandarin)"
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
	q.Question = strings.TrimSpace(q.Question)
	return q
}

func (q Query) Locale() string {
	if q.Field == "med" {
		return "en"
	}
	return ""
}

func (q Query) Country() string {
	if utils.StrContains(q.ProgLang, "nhs", "nice") {
		return "gb"
	}
	return "us"
}

type ContinuousAskQuery struct {
	Query
	ParentId string `json:"parent_id"`
}

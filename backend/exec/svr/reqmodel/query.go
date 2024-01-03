package reqmodel

import (
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

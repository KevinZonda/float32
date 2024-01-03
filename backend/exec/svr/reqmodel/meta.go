package reqmodel

import (
	"encoding/json"
	"github.com/KevinZonda/float32/rag"
)

func NewMeta(searched rag.SearchRawItem) MetaModel {
	var evi []SearchItem
	for _, r := range searched.SpiderResults {
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
		Related:   searched.Related,
	}
}

type MetaModel struct {
	Evidences []SearchItem `json:"evidences"`
	ID        string       `json:"id"`
	Related   []string     `json:"related"`
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

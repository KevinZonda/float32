package main

import "encoding/json"

type MetaModel struct {
	Evidences []SearchItem `json:"evidences"`
}

type SearchItem struct {
	Url   string `json:"url"`
	Title string `json:"title"`
	//Description string `json:"description"`
}

func (m MetaModel) Json() string {
	b, _ := json.Marshal(m)
	return string(b)
}

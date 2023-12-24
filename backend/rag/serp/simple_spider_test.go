package serp

import (
	"encoding/json"
	"testing"
)

func TestSimpleSpider(t *testing.T) {
	s := NewSimpleSpider()
	resp := s.Search("https://blog.kevinzonda.com/post/gan/", "https://go.dev/doc/install")
	jout(resp)
}

func jout(v any) {
	t, _ := json.MarshalIndent(v, "", "  ")
	println(string(t))
}

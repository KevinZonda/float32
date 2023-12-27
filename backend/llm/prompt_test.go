package llm_test

import (
	"fmt"
	"github.com/KevinZonda/float32/llm"
	"testing"
)

func TestPrompt(t *testing.T) {
	ptc := llm.Promptc("", "How to create a file in Go?", "English", "Go", "https://golang.org/pkg/os/#Create")
	fmt.Println(ptc)
}

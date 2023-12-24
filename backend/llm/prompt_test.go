package llm

import (
	"fmt"
	"testing"
)

func TestPrompt(t *testing.T) {
	ptc := Promptc("How to create a file in Go?", "English", "Go", "https://golang.org/pkg/os/#Create")
	fmt.Println(ptc)
}

package serp_test

import (
	"fmt"
	"github.com/KevinZonda/float32/rag/serp"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestGoogle(t *testing.T) {
	godotenv.Load("../../.env")
	fmt.Println(os.Getenv("SERP_DEV"))
	gs := serp.NewGoogleSearch(os.Getenv("SERP_DEV"))
	resp, err := gs.Search("us", "en-us", "How to create a file in Go?")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

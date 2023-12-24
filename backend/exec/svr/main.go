package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/KevinZonda/float32/llm"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
	"strings"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("OPENAI_SB")
	fmt.Println("Work at:", os.Getenv("PWD"))
	fmt.Println("Token found:", token)
	cfg := openai.DefaultConfig(token)
	// cfg.BaseURL = "https://api.openai-sb.com/v1"
	cli := openai.NewClientWithConfig(cfg)
	var history []openai.ChatCompletionMessage
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		question := scanner.Text()
		fmt.Println("Question:", question)
		req := openai.ChatCompletionRequest{
			Temperature:      0.3,
			N:                1,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
			Model:            openai.GPT3Dot5Turbo,
		}

		if len(history) == 0 {
			searched := search("Golang, " + question)
			fmt.Println("Search result:", searched)
			fmt.Println("---------------------------")
			history = append(history, openai.ChatCompletionMessage{
				Content: llm.Promptc(question, "English", "Go", searched),
				Role:    "system",
			})
		}
		req.Messages = history
		//fmt.Println(req)
		resp, err := cli.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			panic(err)
		}
		defer resp.Close()
		respSB := strings.Builder{}
		fmt.Println()
		fmt.Print("<< ")
		for {
			msg, respE := resp.Recv()
			if errors.Is(respE, io.EOF) {
				break
			}
			content := msg.Choices[0].Delta.Content
			respSB.WriteString(content)
			fmt.Print(content)
		}
		history = append(history, openai.ChatCompletionMessage{
			Content: respSB.String(),
			Role:    openai.ChatMessageRoleAssistant,
		})
		fmt.Println()
	}
}

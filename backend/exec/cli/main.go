package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/utils"
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

		content := question
		role := openai.ChatMessageRoleUser

		if len(history) == 0 {
			progLang := "golang"
			if progLang == "general" {
				progLang = ""
			}

			searched := rag.Search(rag.MapProgLang(progLang) + ", " + question)
			fmt.Println("Search result:", searched)
			fmt.Println("---------------------------")
			content = llm.Promptc(question, "English", progLang, searched)
		}
		history = append(history, openai.ChatCompletionMessage{
			Content: content,
			Role:    role,
		})
		req.Messages = history
		//fmt.Println(req)
		resp, err := cli.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Print("<< ")
		respS, _ := recvByLine(resp)
		resp.Close()
		history = append(history, openai.ChatCompletionMessage{
			Content: respS,
			Role:    openai.ChatMessageRoleAssistant,
		})

		fmt.Println()

		req = openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.3,
			N:           1,
			Messages:    llm.Translate("Chinese (Mandarin)", respS),
		}

		resp, err = cli.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			panic(err)
		}
		fmt.Print("<< ")
		respS, _ = recvByLine(resp)
		resp.Close()
		history = append(history, openai.ChatCompletionMessage{
			Content: respS,
			Role:    openai.ChatMessageRoleAssistant,
		})
		fmt.Println()
	}
}

func recvByLine(resp *openai.ChatCompletionStream) (string, error) {
	sb := strings.Builder{}
	buf := strings.Builder{}
	for {
		msg, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		delta := msg.Choices[0].Delta.Content
		delta = utils.CleanStr(delta)
		sb.WriteString(delta)

		if printOutBySubStrs(&sb, &buf, delta, ".", "\n", ";", "。", "？", "?") {
			continue
		}
		buf.WriteString(delta)
	}
	bufS := buf.String()
	if bufS != "" {
		fmt.Print(bufS)
		sb.WriteString(bufS)
	}

	return sb.String(), nil
}

func printOutBySubStr(sb, buf *strings.Builder, delta, subStr string) (needContinue bool) {
	if idx := strings.Index(delta, subStr); idx > 0 {
		toPrint := buf.String() + delta[:idx+1]
		fmt.Print(toPrint)
		buf.Reset()
		sb.WriteString(delta[idx+1:])
		needContinue = true
	}
	return
}

func printOutBySubStrs(sb, buf *strings.Builder, delta string, subStrs ...string) (needContinue bool) {
	for _, subString := range subStrs {
		if printOutBySubStr(sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

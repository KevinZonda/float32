package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
	"strings"
)

var cli *openai.Client

func main() {
	godotenv.Load(".env")
	token := os.Getenv("OPENAI_SB")
	fmt.Println("Work at:", os.Getenv("PWD"))
	fmt.Println("Token found:", token)
	cfg := openai.DefaultConfig(token)
	cli = openai.NewClientWithConfig(cfg)

	g := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	g.Use(cors.New(config))
	g.POST("/query", func(c *gin.Context) {
		var query Query
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		searched := rag.Search(rag.MapProgLang(query.ProgLang) + ", " + query.Question)
		content := llm.Promptc(query.Question, query.Language, query.ProgLang, searched)
		req := openai.ChatCompletionRequest{
			Temperature:      0.3,
			N:                1,
			PresencePenalty:  0,
			FrequencyPenalty: 0,
			Model:            openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Content: content,
					Role:    openai.ChatMessageRoleSystem,
				},
			},
		}

		resp, err := cli.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer resp.Close()

		sb := strings.Builder{}
		buf := strings.Builder{}

		c.Stream(func(w io.Writer) bool {
			msg, err := resp.Recv()
			if errors.Is(err, io.EOF) || err != nil {
				bufS := buf.String()
				if bufS != "" {
					w.Write([]byte(bufS))
					sb.WriteString(bufS)
				}
				return false
			}
			delta := msg.Choices[0].Delta.Content
			delta = utils.CleanStr(delta)
			sb.WriteString(delta)

			if printOutBySubStrs(w, &sb, &buf, delta, "\n", ".", ";", "。", "？", "?") {
				return true
			}
			buf.WriteString(delta)
			return true
		})
		fmt.Println(sb.String())
		log.Println("COUNT:", strings.Count(sb.String(), "\n"))

	})
	g.Run("127.0.0.1:8080")
}

type Query struct {
	Question string `json:"question"`
	ProgLang string `json:"prog_lang"`
	Language string `json:"language"`
}

func printOutBySubStr(w io.Writer, sb, buf *strings.Builder, delta, subStr string) (needContinue bool) {
	if idx := strings.Index(delta, subStr); idx > 0 {
		if subStr == "\n" {
			fmt.Println(idx)
		}
		toPrint := buf.String() + delta[:idx+1]
		w.Write([]byte(toPrint))
		buf.Reset()
		sb.WriteString(delta[idx+1:])
		needContinue = true
	}
	return
}

func printOutBySubStrs(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...string) (needContinue bool) {
	for _, subString := range subStrs {
		if printOutBySubStr(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

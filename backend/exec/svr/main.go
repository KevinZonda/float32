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
	"os"
	"strings"
)

var cli *openai.Client

func initAll() {
	godotenv.Load(".env")
	fmt.Println("Work at:", os.Getenv("PWD"))
	initChatGPT()
	initGin()
}

func initChatGPT() {
	token := os.Getenv("OPENAI_SB")
	cfg := openai.DefaultConfig(token)
	cli = openai.NewClientWithConfig(cfg)
}

var g *gin.Engine

func initGin() {
	g = gin.Default()
	config := cors.DefaultConfig()
	if strings.TrimSpace(os.Getenv("DEBUG")) == "1" {
		gin.SetMode(gin.DebugMode)
		config.AllowAllOrigins = true
		g.Use(cors.New(config))
	} else {
		gin.SetMode(gin.ReleaseMode)
		config.AllowOrigins = []string{"https://float32.app"}
	}
	g.Use(cors.New(config))
}

func startGin() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	g.Run(listenAddr)
}

func main() {
	initAll()

	g.POST("/query", func(c *gin.Context) {
		var query Query
		if err := c.BindJSON(&query); err != nil {
			utils.GinErrorMsg(c, err)
			return
		}
		query = query.Regularize()
		searched := ""
		// TODO: Country fix
		country := "us"
		if utils.StrContains(query.ProgLang, "nhs", "nice") {
			country = "uk"
		}
		searchRaw, err := rag.SearchRaw(country, query.ProgLang+", "+query.Question)
		if err == nil {
			searched = rag.SearchResultsToText(searchRaw)
		}
		// write meta info to Http
		meta := newMeta(searchRaw)
		c.String(200, "%s\r\n", meta.Json())

		content := llm.Promptc(query.Field, query.Question, query.Language, query.ProgLang, searched)
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
				{
					Content: query.Question,
					Role:    openai.ChatMessageRoleUser,
				},
			},
		}

		var resp *openai.ChatCompletionStream

		resp, err = cli.CreateChatCompletionStream(context.Background(), req)

		if err != nil {
			utils.GinErrorMsg(c, err)
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

			if writeBySubStrs(w, &sb, &buf, delta, '\n', '.', ';', '。', '？', '?') {
				return true
			}
			buf.WriteString(delta)
			return true
		})
		fmt.Println(sb.String())
		//log.Println("COUNT:", strings.Count(sb.String(), "\n"))
	})
	startGin()
}

func writeBySubStr(w io.Writer, sb, buf *strings.Builder, delta string, subStr rune) (needContinue bool) {
	rs := []rune(delta)
	if idx := utils.IndexOfRunes(rs, subStr); idx > 0 {
		toPrint := buf.String() + string(rs[:idx+1])
		w.Write([]byte(toPrint))
		buf.Reset()
		sb.WriteString(string(rs[idx+1:]))
		needContinue = true
	}
	return
}

func writeBySubStrs(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if writeBySubStr(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

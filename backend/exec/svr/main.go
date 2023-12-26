package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/rag/serp"
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

func main() {
	godotenv.Load(".env")
	token := os.Getenv("OPENAI_SB")
	listenAddr := os.Getenv("LISTEN_ADDR")

	fmt.Println("Work at:", os.Getenv("PWD"))
	cfg := openai.DefaultConfig(token)
	cli = openai.NewClientWithConfig(cfg)

	g := gin.Default()
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

	g.POST("/query", func(c *gin.Context) {
		var query Query
		err := c.BindJSON(&query)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
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

		//if query.Language != "en" {
		//	var zhResp openai.ChatCompletionResponse
		//	zhResp, err = cli.CreateChatCompletion(context.Background(), req)
		//	if err == nil {
		//		req = openai.ChatCompletionRequest{
		//			Model:       openai.GPT3Dot5Turbo,
		//			Temperature: 0.3,
		//			N:           1,
		//			Messages:    llm.Translate("Chinese", zhResp.Choices[0].Message.Content),
		//		}
		//	} else {
		//		c.JSON(400, gin.H{
		//			"message": err.Error(),
		//		})
		//		return
		//	}
		//}
		resp, err = cli.CreateChatCompletionStream(context.Background(), req)

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

			if printOutBySubStrs(w, &sb, &buf, delta, '\n', '.', ';', '。', '？', '?') {
				return true
			}
			buf.WriteString(delta)
			return true
		})
		fmt.Println(sb.String())
		//log.Println("COUNT:", strings.Count(sb.String(), "\n"))
	})
	g.Run(listenAddr)
}

func printOutBySubStr(w io.Writer, sb, buf *strings.Builder, delta string, subStr rune) (needContinue bool) {
	rs := []rune(delta)
	if idx := idxRunes(rs, subStr); idx > 0 {
		toPrint := buf.String() + string(rs[:idx+1])
		w.Write([]byte(toPrint))
		buf.Reset()
		sb.WriteString(string(rs[idx+1:]))
		needContinue = true
	}
	return
}

func idxRunes(rs []rune, r rune) int {
	for idx, r1 := range rs {
		if r == r1 {
			return idx
		}
	}
	return -1
}

func printOutBySubStrs(w io.Writer, sb, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if printOutBySubStr(w, sb, buf, delta, subString) {
			needContinue = true
			return
		}
	}
	return
}

func newMeta(searched []serp.SpiderResult) MetaModel {
	var evi []SearchItem
	for _, r := range searched {
		if r.Error != nil {
			continue
		}
		evi = append(evi, SearchItem{
			Url:         r.Url,
			Title:       r.Title,
			Description: utils.StrMaxLenSmart(r.Description, 50, "..."),
		})
	}
	return MetaModel{
		Evidences: evi,
	}
}

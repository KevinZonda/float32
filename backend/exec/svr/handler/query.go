package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/exec/svr/dbmodel"
	"github.com/KevinZonda/float32/exec/svr/reqmodel"
	"github.com/KevinZonda/float32/exec/svr/shared"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"strings"
)

func Search(c *gin.Context) {
	var query reqmodel.Query
	if err := c.BindJSON(&query); err != nil {
		utils.GinErrorMsg(c, err)
		return
	}
	query = query.Regularize()
	if query.Question == "" {
		c.String(200, "")
		return
	}

	ans, _ := db.NewAnswer(dbmodel.Answer{
		Question: query.Question,
	})

	searched := ""
	// TODO: Country fix
	country := "us"
	if utils.StrContains(query.ProgLang, "nhs", "nice") {
		country = "gb"
	}
	locale := ""
	if query.Field == "med" {
		locale = "en"
	}
	searchRaw, err := rag.SearchRaw(country, locale, query.ProgLang+", "+query.Question)
	if err == nil {
		searched = rag.SearchResultsToText(searchRaw.SpiderResults)
	}
	// write meta info to Http
	meta := reqmodel.NewMeta(searchRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	bs, _ := json.Marshal(meta.Evidences)
	ans.Evidence = string(bs)

	content := llm.Promptc(query.Field, query.Question, query.Language, query.ProgLang, searched)
	req := openai.ChatCompletionRequest{
		Temperature:      0.15,
		N:                1,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		Model:            openai.GPT3Dot5Turbo1106,
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

	resp, err = shared.Cli.CreateChatCompletionStream(context.Background(), req)

	if err != nil {
		utils.GinErrorMsg(c, errors.New("LLM backend broken"))
		log.Println(err)
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

		if utils.WriteSplitByRune(w, &sb, &buf, delta, '\n', '.', ';', '。', '？', '?') {
			return true
		}
		buf.WriteString(delta)
		return true
	})
	fmt.Println(sb.String(), "\n->", ans)
	ans.FirstAnswer = sb.String()
	ans.IsOk = true
	db.UpdateAnswer(ans)
	//log.Println("COUNT:", strings.Count(sb.String(), "\n"))
}

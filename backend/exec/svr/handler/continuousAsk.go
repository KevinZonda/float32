package handler

import (
	"context"
	"errors"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/exec/svr/dbmodel"
	"github.com/KevinZonda/float32/exec/svr/reqmodel"
	"github.com/KevinZonda/float32/exec/svr/shared"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

func translate(content string) string {
	req := llm.PromptcTranslate(content)
	resp, e := shared.Cli.CreateChatCompletion(context.Background(), req)
	if e != nil {
		log.Println(e)
		return content
	}
	if len(resp.Choices) == 0 {
		return content
	}
	txt := resp.Choices[0].Message.Content
	log.Println("TRANSLATE> ", content, "----->", txt)
	txt = strings.TrimSpace(strings.ReplaceAll(txt, "\n", " "))
	if strings.Contains(strings.ToLower(txt), "err") {
		return content
	}
	return txt
}

func ContinuousAsk(c *gin.Context) {
	var query reqmodel.ContinuousAskQuery
	if err := c.BindJSON(&query); err != nil {
		utils.GinErrorMsg(c, err)
		return
	}

	prevAns, err := db.FindAnswerById(query.ParentId)
	if err != nil {
		utils.GinErrorMsg(c, errors.New("db error"))
		return
	}
	query.Query = query.Regularize()
	if query.Question == "" {
		c.String(200, "")
		return
	}

	ans, _ := db.NewAnswer(dbmodel.Answer{
		Question: query.Question,
	})

	searchRaw, err := rag.SearchRaw(query.Country(), query.Locale(), query.ProgLang+", "+translate(query.Question))
	searched := rag.SearchResultsToText(searchRaw.SpiderResults)

	// write meta info to Http
	meta := reqmodel.NewMeta(searchRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	ans.Evidence = utils.Json(meta.Evidences)

	content := llm.Promptc(query.Language, query.Field, query.Question, query.ProgLang, searched)

	req := openai.ChatCompletionRequest{
		Temperature: 0.15,
		N:           1,
		Model:       openai.GPT3Dot5Turbo1106,
		Messages: []openai.ChatCompletionMessage{
			{
				Content: content,
				Role:    openai.ChatMessageRoleSystem,
			},
			{
				Content: prevAns.Question,
				Role:    openai.ChatMessageRoleUser,
			},
			{
				Content: prevAns.FirstAnswer,
				Role:    openai.ChatMessageRoleAssistant,
			},
			{
				Content: query.Question,
				Role:    openai.ChatMessageRoleUser,
			},
		},
	}

	ans.FirstAnswer = chatStreamToGin(c, req)
	ans.IsOk = true
	db.UpdateAnswer(ans)
}

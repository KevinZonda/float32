package handler

import (
	"errors"
	"github.com/KevinZonda/float32/exec/svr/db"
	"github.com/KevinZonda/float32/exec/svr/dbmodel"
	"github.com/KevinZonda/float32/exec/svr/reqmodel"
	"github.com/KevinZonda/float32/llm"
	"github.com/KevinZonda/float32/rag"
	"github.com/KevinZonda/float32/utils"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

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

	searchRaw, err := rag.SearchRaw(query.Country(), query.Locale(), query.ProgLang+", "+query.Question)
	searched := rag.SearchResultsToText(searchRaw.SpiderResults)

	// write meta info to Http
	meta := reqmodel.NewMeta(searchRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	ans.Evidence = utils.Json(meta.Evidences)

	content := llm.Promptc(query.Language, query.Question, query.Field, query.ProgLang, searched)

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

	chatStreamToGin(c, req)
}

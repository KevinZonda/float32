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
	searchRaw, err := rag.SearchRaw(query.Country(), query.Locale(), query.ProgLang+", "+translate(query.Question))
	if err == nil {
		searched = rag.SearchResultsToText(searchRaw.SpiderResults)
	}
	// write meta info to Http
	meta := reqmodel.NewMeta(searchRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	ans.Evidence = utils.Json(meta.Evidences)

	content := llm.Promptc(query.Language, query.Field, query.Question, query.ProgLang, searched)

	req := utils.ModelGPT35Request([]openai.ChatCompletionMessage{
		utils.ChatMsgFromSystem(content),
		utils.ChatMsgFromUser(query.Question),
	})

	ans.FirstAnswer = chatStreamToGin(c, req)
	ans.IsOk = true
	db.UpdateAnswer(ans)
	//log.Println("COUNT:", strings.Count(sb.String(), "\n"))
}

func chatStreamToGin(c *gin.Context, req openai.ChatCompletionRequest) (completeAns string) {
	var resp *openai.ChatCompletionStream

	resp, err := shared.Cli.CreateChatCompletionStream(context.Background(), req)

	if err != nil {
		utils.GinErrorMsgTxt(c, "LLM backend broken")
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

		if utils.WriteSplitByRune(w, &buf, delta, '\n', '.', ';', '。', '？', '?') {
			return true
		}
		return true
	})
	completeAns = sb.String()
	return
}

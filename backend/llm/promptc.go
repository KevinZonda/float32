package llm

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

func systemPrompt(ptsName Field, varMap map[string]string) string {
	compiled := _pts[ptsName].CompileWithOption(varMap, false)
	return strings.TrimSpace(compiled.Prompts[0].Prompt)
}

func Promptc(lang string, field string, question string, guide string, context any) string {
	ptsField := CodeField
	switch field {

	case "med":
		ptsField = MedField
	default:
		ptsField = CodeField
		if guide != "" {
			guide = fmt.Sprintf(" When it comes to answers in code, please express them in the %s programming language.", guide)
		} else {
			guide = " Please stand in the perspective of a programmer or an advanced software engineer. When it comes to answers in code, please express them in the Java programming language except user specified any programming language."
		}
	}
	varMap := map[string]string{
		"lang":     lang,
		"guide":    guide,
		"question": question,
		"context":  fmt.Sprint(context),
	}
	return systemPrompt(ptsField, varMap)
}

func PromptcTranslate(content string) openai.ChatCompletionRequest {
	varMap := map[string]string{
		"content": content,
	}
	compiled := _pts[TrannslateField].CompileWithOption(varMap, false)
	return openai.ChatCompletionRequest{
		Temperature: 0.00000000001,
		Model:       openai.GPT3Dot5Turbo1106,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: strings.TrimSpace(compiled.Prompts[0].Prompt),
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "<NeedTranslate>\n" + content + "\n</NeedTranslate>",
			},
		},
	}

}

type Field string

const (
	CodeField       Field = "code"
	MedField        Field = "med"
	TrannslateField Field = "translate"
)

var _pts map[Field]*prompt.PromptC

func init() {
	log.Println("Loading promptc...")
	_pts = map[Field]*prompt.PromptC{
		CodeField:       loadPromptc("code.promptc"),
		MedField:        loadPromptc("med.promptc"),
		TrannslateField: loadPromptc("translate.promptc"),
	}
}

func loadPromptc(path string) *prompt.PromptC {
	pt, err := iox.ReadAllText(path)
	if err != nil {
		log.Println("Failed to load promptc", err)
		panic(err)
	}
	return prompt.ParsePromptC(pt)
}

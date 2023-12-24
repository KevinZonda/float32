package llm

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	"github.com/sashabaranov/go-openai"
)

func Promptc(question string, answerIn string, programmingLanguage string, searchResult any) string {
	if programmingLanguage != "" {
		programmingLanguage = fmt.Sprintf(" When it comes to answers in code, please express them in the %s programming language.", programmingLanguage)
	} else {
		programmingLanguage = " Please stand in the perspective of a programmer or an advanced software engineer. When it comes to answers in code, please express them in the Java programming language except user specified any programming language."
	}
	if answerIn != "" {
		answerIn = "English"
	}
	varMap := map[string]string{
		"lang":        answerIn,
		"programLang": programmingLanguage,
		"question":    question,
		"query":       fmt.Sprint(searchResult),
	}

	compiled := _ptc.CompileWithOption(varMap, false)
	return compiled.Prompts[0].Prompt
}

func Translate(toLang string, content string) []openai.ChatCompletionMessage {
	varMap := map[string]string{
		"to":      toLang,
		"content": content,
	}

	compiled := _trans.CompileWithOption(varMap, false)
	return []openai.ChatCompletionMessage{
		{
			Content: compiled.Prompts[0].Prompt,
			Role:    openai.ChatMessageRoleAssistant,
		},
		{
			Content: compiled.Prompts[1].Prompt,
			Role:    openai.ChatMessageRoleUser,
		},
	}
}

var _ptc *prompt.PromptC
var _trans *prompt.PromptC

func init() {
	pt, err := iox.ReadAllText("prompt.promptc")
	if err != nil {
		panic(err)
	}

	_ptc = prompt.ParsePromptC(pt)

	pt, err = iox.ReadAllText("translate.promptc")
	if err != nil {
		panic(err)
	}

	_trans = prompt.ParsePromptC(pt)
}

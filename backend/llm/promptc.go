package llm

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	"github.com/sashabaranov/go-openai"
)

func Promptc(question string, answerIn string, programmingLanguage string, searchResult any) string {
	if answerIn == "简体中文" {
		return promptcZh(question, answerIn, programmingLanguage, searchResult)
	}
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

func promptcZh(question string, answerIn string, programmingLanguage string, searchResult any) string {
	if programmingLanguage != "" {
		programmingLanguage = fmt.Sprintf("如果需要用代码作答，请用 %s 程序语言来表达。", programmingLanguage)
	} else {
		programmingLanguage = "请站在程序员或高级软件工程师的角度作答。用代码作答时，除用户指定的任何编程语言外，请用 Java 编程语言表达。"
	}
	if answerIn != "" {
		answerIn = "简体中文"
	}
	varMap := map[string]string{
		"lang":        answerIn,
		"programLang": programmingLanguage,
		"question":    question,
		"query":       fmt.Sprint(searchResult),
	}

	compiled := _ptc_zh.CompileWithOption(varMap, false)
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
var _ptc_zh *prompt.PromptC

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

	pt, err = iox.ReadAllText("prompt_zh.promptc")
	if err != nil {
		panic(err)
	}

	_ptc_zh = prompt.ParsePromptC(pt)
}

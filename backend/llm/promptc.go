package llm

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	"github.com/sashabaranov/go-openai"
)

func Promptc(field string, question string, answerIn string, programmingLanguage string, searchResult any) string {
	if answerIn == "简体中文" {
		return promptcZh(field, question, answerIn, programmingLanguage, searchResult)
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

func promptcZh(field string, question string, answerIn string, programmingLanguage string, searchResult any) string {
	if field == "med" {
		varMap := map[string]string{
			"lang":        answerIn,
			"programLang": programmingLanguage,
			"question":    question,
			"query":       fmt.Sprint(searchResult),
		}

		compiled := _med_zh.CompileWithOption(varMap, false)
		return compiled.Prompts[0].Prompt
	}

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
var _med_zh *prompt.PromptC

func init() {
	fmt.Println("Loading promptc...")
	_ptc = loadPromptc("prompt.promptc")
	_trans = loadPromptc("translate.promptc")
	_ptc_zh = loadPromptc("prompt_zh.promptc")
	_med_zh = loadPromptc("med_prompt_zh.promptc")
}

func loadPromptc(path string) *prompt.PromptC {
	pt, err := iox.ReadAllText(path)
	if err != nil {
		fmt.Println("Failed to load promptc", err)
		panic(err)
	}
	return prompt.ParsePromptC(pt)
}

package llm

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
	"github.com/sashabaranov/go-openai"
)

func Promptc(field string, question string, answerIn string, guide string, context any) string {
	if answerIn == "简体中文" {
		return promptcZh(field, question, guide, context)
	}
	return promptc(field, question, guide, context)
}

func firstPromptStr(ptsName Field, varMap map[string]string) string {
	compiled := _pts[ptsName].CompileWithOption(varMap, false)
	return compiled.Prompts[0].Prompt
}

func promptc(field string, question string, guide string, context any) string {
	ptsField := CodeField
	switch field {
	default:
		ptsField = CodeField
		if guide != "" {
			guide = fmt.Sprintf(" When it comes to answers in code, please express them in the %s programming language.", guide)
		} else {
			guide = " Please stand in the perspective of a programmer or an advanced software engineer. When it comes to answers in code, please express them in the Java programming language except user specified any programming language."
		}

	}
	varMap := map[string]string{
		"lang":     "English",
		"guide":    guide,
		"question": question,
		"context":  fmt.Sprint(context),
	}
	return firstPromptStr(ptsField, varMap)
}

func promptcZh(field string, question string, guide string, context any) string {
	ptsField := CodeZhField

	switch field {
	case "med":
		ptsField = MedZhField
	default: // "code"
		ptsField = CodeZhField
		if guide != "" {
			guide = fmt.Sprintf("如果需要用代码作答，请用 %s 程序语言来表达。", guide)
		} else {
			guide = "请站在程序员或高级软件工程师的角度作答。用代码作答时，除用户指定的任何编程语言外，请用 Java 编程语言表达。"
		}
	}

	varMap := map[string]string{
		"lang":     "简体中文",
		"guide":    guide,
		"question": question,
		"context":  fmt.Sprint(context),
	}
	return firstPromptStr(ptsField, varMap)
}

func Translate(toLang string, content string) []openai.ChatCompletionMessage {
	varMap := map[string]string{
		"to":      toLang,
		"content": content,
	}

	compiled := _pts[TranslationField].CompileWithOption(varMap, false)
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

type Field string

const (
	CodeField        Field = "code"
	MedZhField       Field = "med_zh"
	CodeZhField      Field = "code_zh"
	TranslationField Field = "trans"
)

var _pts map[Field]*prompt.PromptC

func init() {
	fmt.Println("Loading promptc...")
	_pts = make(map[Field]*prompt.PromptC)
	_pts[CodeField] = loadPromptc("prompt.promptc")
	_pts[TranslationField] = loadPromptc("translate.promptc")
	_pts[CodeZhField] = loadPromptc("prompt_zh.promptc")
	_pts[MedZhField] = loadPromptc("med_prompt_zh.promptc")
}

func loadPromptc(path string) *prompt.PromptC {
	pt, err := iox.ReadAllText(path)
	if err != nil {
		fmt.Println("Failed to load promptc", err)
		panic(err)
	}
	return prompt.ParsePromptC(pt)
}

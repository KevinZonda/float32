package utils

import "github.com/sashabaranov/go-openai"

func ChatMsgFromUser(txt string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: txt,
	}
}

func ChatMsgFromSystem(txt string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: txt,
	}
}

func ChatMsgFromAssistant(txt string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: txt,
	}
}

func ModelGPT35Request(msg []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Temperature: 0.1,
		N:           1,
		Model:       openai.GPT3Dot5Turbo0125,
		Messages:    msg,
	}
}
